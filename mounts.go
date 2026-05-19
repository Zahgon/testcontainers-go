package testcontainers

import (
	"errors"
)

const (
	MountTypeBind MountType = iota // Deprecated: Use MountTypeVolume instead
	MountTypeVolume
	MountTypeTmpfs
	MountTypePipe
	MountTypeImage
)

var (
	ErrDuplicateMountTarget = errors.New("duplicate mount target detected")
	ErrInvalidBindMount     = errors.New("invalid bind mount")
)

var (
	_ ContainerMountSource = (*GenericBindMountSource)(nil) // Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
	_ ContainerMountSource = (*GenericVolumeMountSource)(nil)
	_ ContainerMountSource = (*GenericTmpfsMountSource)(nil)
	_ ContainerMountSource = (*GenericImageMountSource)(nil)
)

type (
	// ContainerMounts represents a collection of mounts for a container
	ContainerMounts []ContainerMount
	MountType       uint
)

// ContainerMountSource is the base for all mount sources
type ContainerMountSource interface {
	// Source will be used as Source field in the final mount
	// this might either be a volume name, a host path or might be empty e.g. for Tmpfs
	Source() string

	// Type determines the final mount type
	// possible options are limited by the Docker API
	Type() MountType
}

// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
// GenericBindMountSource implements ContainerMountSource and represents a bind mount
// Optionally mount.BindOptions might be added for advanced scenarios
type GenericBindMountSource struct {
	// HostPath is the path mounted into the container
	// the same host path might be mounted to multiple locations within a single container
	HostPath string
}

// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
func (s GenericBindMountSource) Source() string {
	_ = "STUB: not implemented"

	// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
	return ""
}

func (GenericBindMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// GenericVolumeMountSource implements ContainerMountSource and represents a volume mount
	new(MountType)
}

type GenericVolumeMountSource struct {
	// Name refers to the name of the volume to be mounted
	// the same volume might be mounted to multiple locations within a single container
	Name string
}

func (s GenericVolumeMountSource) Source() string { _ = "STUB: not implemented"; return "" }

func (GenericVolumeMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// GenericTmpfsMountSource implements ContainerMountSource and represents a TmpFS mount
	// Optionally mount.TmpfsOptions might be added for advanced scenarios
	new(MountType)
}

type GenericTmpfsMountSource struct{}

func (s GenericTmpfsMountSource) Source() string { _ = "STUB: not implemented"; return "" }

func (GenericTmpfsMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// ContainerMountTarget represents the target path within a container where the mount will be available
	// Note that mount targets must be unique. It's not supported to mount different sources to the same target.
	new(MountType)
}

type ContainerMountTarget string

func (t ContainerMountTarget) Target() string {
	_ = "STUB: not implemented"

	// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
	// BindMount returns a new ContainerMount with a GenericBindMountSource as source
	// This is a convenience method to cover typical use cases.
	return ""
}

func BindMount(hostPath string, mountTarget ContainerMountTarget) ContainerMount {
	_ = "STUB: not implemented"
	return *new(ContainerMount)
}

// VolumeMount returns a new ContainerMount with a GenericVolumeMountSource as source
// This is a convenience method to cover typical use cases.
func VolumeMount(volumeName string, mountTarget ContainerMountTarget) ContainerMount {
	_ = "STUB: not implemented"
	return *new(ContainerMount)
}

// ImageMount returns a new ContainerMount with a GenericImageMountSource as source
// This is a convenience method to cover typical use cases.
func ImageMount(imageName string, subpath string, mountTarget ContainerMountTarget) ContainerMount {
	_ = "STUB: not implemented"
	return *new(ContainerMount)
}

// Mounts returns a ContainerMounts to support a more fluent API
func Mounts(mounts ...ContainerMount) ContainerMounts {
	_ = "STUB: not implemented"

	// ContainerMount models a mount into a container
	return *new(ContainerMounts)
}

type ContainerMount struct {
	// Source is typically either a GenericVolumeMountSource, as BindMount is not supported by all Docker environments
	Source ContainerMountSource
	// Target is the path where the mount should be mounted within the container
	Target ContainerMountTarget
	// ReadOnly determines if the mount should be read-only
	ReadOnly bool
}

// GenericImageMountSource implements ContainerMountSource and represents an image mount
type GenericImageMountSource struct {
	// imageName refers to the name of the image to be mounted
	// the same image might be mounted to multiple locations within a single container
	imageName string
	// subpath is the path within the image to be mounted
	subpath string
}

// NewGenericImageMountSource creates a new GenericImageMountSource
func NewGenericImageMountSource(imageName string, subpath string) GenericImageMountSource {
	_ = "STUB: not implemented"
	return *new(GenericImageMountSource)
}

// Source returns the name of the image to be mounted
func (s GenericImageMountSource) Source() string {
	_ = "STUB: not implemented"

	// Type returns the type of the mount
	return ""
}

func (GenericImageMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// Validate validates the source of the mount
	new(MountType)
}

func (s GenericImageMountSource) Validate() error { _ = "STUB: not implemented"; return nil }
