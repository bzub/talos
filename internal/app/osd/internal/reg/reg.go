// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package reg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"

	criconstants "github.com/containerd/cri/pkg/constants"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prometheus/procfs"
	"google.golang.org/grpc"

	"github.com/talos-systems/talos/api/common"
	osapi "github.com/talos-systems/talos/api/os"
	"github.com/talos-systems/talos/internal/pkg/containers"
	"github.com/talos-systems/talos/internal/pkg/containers/containerd"
	"github.com/talos-systems/talos/internal/pkg/containers/cri"
	"github.com/talos-systems/talos/internal/pkg/kmsg"
	"github.com/talos-systems/talos/pkg/constants"
)

// Registrator is the concrete type that implements the factory.Registrator and
// osapi.OSDServer interfaces.
type Registrator struct{}

// Register implements the factory.Registrator interface.
func (r *Registrator) Register(s *grpc.Server) {
	osapi.RegisterOSServiceServer(s, r)
}

// Containers implements the osapi.OSDServer interface.
func (r *Registrator) Containers(ctx context.Context, in *osapi.ContainersRequest) (reply *osapi.ContainersResponse, err error) {
	inspector, err := getContainerInspector(ctx, in.Namespace, in.Driver)
	if err != nil {
		return nil, err
	}
	// nolint: errcheck
	defer inspector.Close()

	pods, err := inspector.Pods()
	if err != nil {
		// fatal error
		if pods == nil {
			return nil, err
		}
		// TODO: only some failed, need to handle it better via client
		log.Println(err.Error())
	}

	containers := []*osapi.ContainerInfo{}

	for _, pod := range pods {
		for _, container := range pod.Containers {
			container := &osapi.ContainerInfo{
				Namespace: in.Namespace,
				Id:        container.Display,
				PodId:     pod.Name,
				Name:      container.Name,
				Image:     container.Image,
				Pid:       container.Pid,
				Status:    container.Status,
			}
			containers = append(containers, container)
		}
	}

	reply = &osapi.ContainersResponse{
		Messages: []*osapi.Container{
			{
				Containers: containers,
			},
		},
	}

	return reply, nil
}

// Stats implements the osapi.OSDServer interface.
// nolint: gocyclo
func (r *Registrator) Stats(ctx context.Context, in *osapi.StatsRequest) (reply *osapi.StatsResponse, err error) {
	inspector, err := getContainerInspector(ctx, in.Namespace, in.Driver)
	if err != nil {
		return nil, err
	}
	// nolint: errcheck
	defer inspector.Close()

	pods, err := inspector.Pods()
	if err != nil {
		// fatal error
		if pods == nil {
			return nil, err
		}
		// TODO: only some failed, need to handle it better via client
		log.Println(err.Error())
	}

	stats := []*osapi.Stat{}

	for _, pod := range pods {
		for _, container := range pod.Containers {
			if container.Metrics == nil {
				continue
			}

			stat := &osapi.Stat{
				Namespace:   in.Namespace,
				Id:          container.Display,
				PodId:       pod.Name,
				Name:        container.Name,
				MemoryUsage: container.Metrics.MemoryUsage,
				CpuUsage:    container.Metrics.CPUUsage,
			}

			stats = append(stats, stat)
		}
	}

	reply = &osapi.StatsResponse{
		Messages: []*osapi.Stats{
			{
				Stats: stats,
			},
		},
	}

	return reply, nil
}

// Restart implements the osapi.OSDServer interface.
func (r *Registrator) Restart(ctx context.Context, in *osapi.RestartRequest) (*osapi.RestartResponse, error) {
	inspector, err := getContainerInspector(ctx, in.Namespace, in.Driver)
	if err != nil {
		return nil, err
	}
	// nolint: errcheck
	defer inspector.Close()

	container, err := inspector.Container(in.Id)
	if err != nil {
		return nil, err
	}

	if container == nil {
		return nil, fmt.Errorf("container %q not found", in.Id)
	}

	err = container.Kill(syscall.SIGTERM)
	if err != nil {
		return nil, err
	}

	return &osapi.RestartResponse{}, nil
}

// Dmesg implements the osapi.OSDServer interface.
//
//nolint: gocyclo
func (r *Registrator) Dmesg(req *osapi.DmesgRequest, srv osapi.OSService_DmesgServer) error {
	ctx := srv.Context()

	var options []kmsg.Option

	if req.Follow {
		options = append(options, kmsg.Follow())
	}

	if req.Tail {
		options = append(options, kmsg.FromTail())
	}

	reader, err := kmsg.NewReader(options...)
	if err != nil {
		return fmt.Errorf("error opening /dev/kmsg reader: %w", err)
	}
	defer reader.Close() //nolint: errcheck

	ch := reader.Scan(ctx)

	for {
		select {
		case <-ctx.Done():
			if err = reader.Close(); err != nil {
				return err
			}
		case packet, ok := <-ch:
			if !ok {
				return nil
			}

			if packet.Err != nil {
				err = srv.Send(&common.Data{
					Metadata: &common.Metadata{
						Error: packet.Err.Error(),
					},
				})
			} else {
				msg := packet.Message
				err = srv.Send(&common.Data{
					Bytes: []byte(fmt.Sprintf("%s: %7s: [%s]: %s", msg.Facility, msg.Priority, msg.Timestamp.Format(time.RFC3339Nano), msg.Message)),
				})
			}

			if err != nil {
				return err
			}
		}
	}
}

// Processes implements the osapi.OSDServer interface
func (r *Registrator) Processes(ctx context.Context, in *empty.Empty) (reply *osapi.ProcessesResponse, err error) {
	procs, err := procfs.AllProcs()
	if err != nil {
		return nil, err
	}

	processes := make([]*osapi.ProcessInfo, 0, len(procs))

	var (
		command    string
		executable string
		args       []string
		stats      procfs.ProcStat
	)

	for _, proc := range procs {
		command, err = proc.Comm()
		if err != nil {
			return nil, err
		}

		executable, err = proc.Executable()
		if err != nil {
			return nil, err
		}

		args, err = proc.CmdLine()
		if err != nil {
			return nil, err
		}

		stats, err = proc.Stat()
		if err != nil {
			return nil, err
		}

		p := &osapi.ProcessInfo{
			Pid:            int32(proc.PID),
			Ppid:           int32(stats.PPID),
			State:          stats.State,
			Threads:        int32(stats.NumThreads),
			CpuTime:        stats.CPUTime(),
			VirtualMemory:  uint64(stats.VirtualMemory()),
			ResidentMemory: uint64(stats.ResidentMemory()),
			Command:        command,
			Executable:     executable,
			Args:           strings.Join(args, " "),
		}

		processes = append(processes, p)
	}

	reply = &osapi.ProcessesResponse{
		Messages: []*osapi.Process{
			{
				Processes: processes,
			},
		},
	}

	return reply, nil
}

func getContainerInspector(ctx context.Context, namespace string, driver common.ContainerDriver) (containers.Inspector, error) {
	switch driver {
	case common.ContainerDriver_CRI:
		if namespace != criconstants.K8sContainerdNamespace {
			return nil, errors.New("CRI inspector is supported only for K8s namespace")
		}

		return cri.NewInspector(ctx)
	case common.ContainerDriver_CONTAINERD:
		addr := constants.ContainerdAddress
		if namespace == constants.SystemContainerdNamespace {
			addr = constants.SystemContainerdAddress
		}

		return containerd.NewInspector(ctx, namespace, containerd.WithContainerdAddress(addr))
	default:
		return nil, fmt.Errorf("unsupported driver %q", driver)
	}
}

// Memory implements the osdapi.OSDServer interface.
func (r *Registrator) Memory(ctx context.Context, in *empty.Empty) (reply *osapi.MemoryResponse, err error) {
	proc, err := procfs.NewDefaultFS()
	if err != nil {
		return nil, err
	}

	info, err := proc.Meminfo()
	if err != nil {
		return nil, err
	}

	meminfo := &osapi.MemInfo{
		Memtotal:          info.MemTotal,
		Memfree:           info.MemFree,
		Memavailable:      info.MemAvailable,
		Buffers:           info.Buffers,
		Cached:            info.Cached,
		Swapcached:        info.SwapCached,
		Active:            info.Active,
		Inactive:          info.Inactive,
		Activeanon:        info.ActiveAnon,
		Inactiveanon:      info.InactiveAnon,
		Activefile:        info.ActiveFile,
		Inactivefile:      info.InactiveFile,
		Unevictable:       info.Unevictable,
		Mlocked:           info.Mlocked,
		Swaptotal:         info.SwapTotal,
		Swapfree:          info.SwapFree,
		Dirty:             info.Dirty,
		Writeback:         info.Writeback,
		Anonpages:         info.AnonPages,
		Mapped:            info.Mapped,
		Shmem:             info.Shmem,
		Slab:              info.Slab,
		Sreclaimable:      info.SReclaimable,
		Sunreclaim:        info.SUnreclaim,
		Kernelstack:       info.KernelStack,
		Pagetables:        info.PageTables,
		Nfsunstable:       info.NFSUnstable,
		Bounce:            info.Bounce,
		Writebacktmp:      info.WritebackTmp,
		Commitlimit:       info.CommitLimit,
		Committedas:       info.CommittedAS,
		Vmalloctotal:      info.VmallocTotal,
		Vmallocused:       info.VmallocUsed,
		Vmallocchunk:      info.VmallocChunk,
		Hardwarecorrupted: info.HardwareCorrupted,
		Anonhugepages:     info.AnonHugePages,
		Shmemhugepages:    info.ShmemHugePages,
		Shmempmdmapped:    info.ShmemPmdMapped,
		Cmatotal:          info.CmaTotal,
		Cmafree:           info.CmaFree,
		Hugepagestotal:    info.HugePagesTotal,
		Hugepagesfree:     info.HugePagesFree,
		Hugepagesrsvd:     info.HugePagesRsvd,
		Hugepagessurp:     info.HugePagesSurp,
		Hugepagesize:      info.Hugepagesize,
		Directmap4K:       info.DirectMap4k,
		Directmap2M:       info.DirectMap2M,
		Directmap1G:       info.DirectMap1G,
	}

	reply = &osapi.MemoryResponse{
		Messages: []*osapi.Memory{
			{
				Meminfo: meminfo,
			},
		},
	}

	return reply, err
}
