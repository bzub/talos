package service

import (
	"fmt"
	"io/ioutil"

	"github.com/autonomy/dianemo/initramfs/src/init/pkg/service/conditions"
	"github.com/autonomy/dianemo/initramfs/src/init/pkg/userdata"
)

const crioConf = `
# The "crio" table contains all of the server options.
[crio]

# root is a path to the "root directory". CRIO stores all of its data,
# including container images, in this directory.
root = "/var/lib/containers/storage"

# run is a path to the "run directory". CRIO stores all of its state
# in this directory.
runroot = "/var/run/containers/storage"

# storage_driver select which storage driver is used to manage storage
# of images and containers.
storage_driver = "overlay"

# storage_option is used to pass an option to the storage driver.
storage_option = [
]

# The "crio.api" table contains settings for the kubelet/gRPC interface.
[crio.api]

# listen is the path to the AF_LOCAL socket on which crio will listen.
listen = "/var/run/crio/crio.sock"

# stream_address is the IP address on which the stream server will listen
stream_address = ""

# stream_port is the port on which the stream server will listen
stream_port = "10010"

# file_locking is whether file-based locking will be used instead of
# in-memory locking
file_locking = true

# The "crio.runtime" table contains settings pertaining to the OCI
# runtime used and options for how to set up and manage the OCI runtime.
[crio.runtime]

# runtime is the OCI compatible runtime used for trusted container workloads.
# This is a mandatory setting as this runtime will be the default one
# and will also be used for untrusted container workloads if
# runtime_untrusted_workload is not set.
runtime = "/bin/runc"

# runtime_untrusted_workload is the OCI compatible runtime used for untrusted
# container workloads. This is an optional setting, except if
# default_container_trust is set to "untrusted".
runtime_untrusted_workload = ""

# default_workload_trust is the default level of trust crio puts in container
# workloads. It can either be "trusted" or "untrusted", and the default
# is "trusted".
# Containers can be run through different container runtimes, depending on
# the trust hints we receive from kubelet:
# - If kubelet tags a container workload as untrusted, crio will try first to
# run it through the untrusted container workload runtime. If it is not set,
# crio will use the trusted runtime.
# - If kubelet does not provide any information about the container workload trust
# level, the selected runtime will depend on the default_container_trust setting.
# If it is set to "untrusted", then all containers except for the host privileged
# ones, will be run by the runtime_untrusted_workload runtime. Host privileged
# containers are by definition trusted and will always use the trusted container
# runtime. If default_container_trust is set to "trusted", crio will use the trusted
# container runtime for all containers.
default_workload_trust = "trusted"

# no_pivot instructs the runtime to not use pivot_root, but instead use MS_MOVE
no_pivot = false

# conmon is the path to conmon binary, used for managing the runtime.
conmon = "/usr/local/libexec/crio/conmon"

# conmon_env is the environment variable list for conmon process,
# used for passing necessary environment variable to conmon or runtime.
conmon_env = [
	"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
]

# selinux indicates whether or not SELinux will be used for pod
# separation on the host. If you enable this flag, SELinux must be running
# on the host.
selinux = false

# seccomp_profile is the seccomp json profile path which is used as the
# default for the runtime.
seccomp_profile = "/etc/crio/seccomp.json"

# apparmor_profile is the apparmor profile name which is used as the
# default for the runtime.
apparmor_profile = "crio-default"

# cgroup_manager is the cgroup management implementation to be used
# for the runtime.
cgroup_manager = "cgroupfs"

# hooks_dir_path is the oci hooks directory for automatically executed hooks
hooks_dir_path = "/var/containers/oci/hooks.d"

# default_mounts is the mounts list to be mounted for the container when created
default_mounts = [
]

# pids_limit is the number of processes allowed in a container
pids_limit = 1024

# enable using a shared PID namespace for containers in a pod
enable_shared_pid_namespace = false

# log_size_max is the max limit for the container log size in bytes.
# Negative values indicate that no limit is imposed.
log_size_max = 1000000

# The "crio.image" table contains settings pertaining to the
# management of OCI images.
[crio.image]

# default_transport is the prefix we try prepending to an image name if the
# image name as we receive it can't be parsed as a valid source reference
default_transport = "docker://"

# pause_image is the image which we use to instantiate infra containers.
pause_image = "kubernetes/pause"

# pause_command is the command to run in a pause_image to have a container just
# sit there.  If the image contains the necessary information, this value need
# not be specified.
pause_command = "/pause"

# signature_policy is the name of the file which decides what sort of policy we
# use when deciding whether or not to trust an image that we've pulled.
# Outside of testing situations, it is strongly advised that this be left
# unspecified so that the default system-wide policy will be used.
signature_policy = ""

# image_volumes controls how image volumes are handled.
# The valid values are mkdir and ignore.
image_volumes = "mkdir"

# insecure_registries is used to skip TLS verification when pulling images.
insecure_registries = [
]

# registries is used to specify a comma separated list of registries to be used
# when pulling an unqualified image (e.g. fedora:rawhide).
registries = [
    "docker.io",
]

# The "crio.network" table contains settings pertaining to the
# management of CNI plugins.
[crio.network]

# network_dir is is where CNI network configuration
# files are stored.
network_dir = "/etc/cni/net.d/"

# plugin_dir is is where CNI plugin binaries are stored.
plugin_dir = "/opt/cni/bin/"
`

const crioPolicy = `
{
    "default": [
        {
            "type": "insecureAcceptAnything"
        }
    ]
}
`

// CRIO implements the Service interface. It serves as the concrete type with
// the required methods.
type CRIO struct{}

// Pre implements the Service interface.
func (p *CRIO) Pre(data userdata.UserData) error {
	if err := ioutil.WriteFile("/etc/crio/crio.conf", []byte(crioConf), 0644); err != nil {
		return fmt.Errorf("write crio.conf: %s", err.Error())
	}
	if err := ioutil.WriteFile("/etc/containers/policy.json", []byte(crioPolicy), 0644); err != nil {
		return fmt.Errorf("write policy.json: %s", err.Error())
	}

	return nil
}

// Cmd implements the Service interface.
func (p *CRIO) Cmd(data userdata.UserData) (name string, args []string) {
	name = "/bin/crio"
	args = []string{}

	return name, args
}

// Condition implements the Service interface.
func (p *CRIO) Condition(data userdata.UserData) func() (bool, error) {
	return conditions.None()
}

// Env implements the Service interface.
func (p *CRIO) Env() []string { return []string{} }

// Type implements the Service interface.
func (p *CRIO) Type() Type { return Forever }
