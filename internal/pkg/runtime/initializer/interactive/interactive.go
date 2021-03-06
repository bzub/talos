// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interactive

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"

	"github.com/talos-systems/go-procfs/procfs"

	"github.com/talos-systems/talos/cmd/installer/pkg"
	"github.com/talos-systems/talos/internal/pkg/runtime"
	"github.com/talos-systems/talos/pkg/blockdevice/probe"
	"github.com/talos-systems/talos/pkg/constants"
)

// Interactive is an initializer that performs an installation by prompting the
// user.
type Interactive struct{}

// Initialize implements the Initializer interface.
func (i *Interactive) Initialize(r runtime.Runtime) (err error) {
	var dev *probe.ProbedBlockDevice

	dev, err = probe.GetDevWithFileSystemLabel(constants.ISOFilesystemLabel)
	if err != nil {
		return fmt.Errorf("failed to find %s iso: %w", constants.ISOFilesystemLabel, err)
	}

	// nolint: errcheck
	defer dev.Close()

	if err = unix.Mount(dev.Path, "/tmp", dev.SuperBlock.Type(), unix.MS_RDONLY, ""); err != nil {
		return err
	}

	for _, f := range []string{"/tmp/usr/install/vmlinuz", "/tmp/usr/install/initramfs.xz"} {
		var source []byte

		source, err = ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile("/"+filepath.Base(f), source, 0644); err != nil {
			return err
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Talos configuration URL: ")

	endpoint, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	cmdline := procfs.NewDefaultCmdline()
	cmdline.Append("initrd", filepath.Join("/", "default", constants.InitramfsAsset))
	cmdline.Append(constants.KernelParamPlatform, r.Platform().Name())
	cmdline.Append(constants.KernelParamConfig, endpoint)

	var installer *pkg.Installer

	installer, err = pkg.NewInstaller(cmdline, r.Config().Machine().Install())
	if err != nil {
		return err
	}

	if err = installer.Install(r.Sequence()); err != nil {
		return fmt.Errorf("failed to install: %w", err)
	}

	return unix.Reboot(int(unix.LINUX_REBOOT_CMD_RESTART))
}
