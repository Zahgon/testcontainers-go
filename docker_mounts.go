package testcontainers

import (
	"github.com/moby/moby/api/types/mount"
)

var mountTypeMapping = map[MountType]mount.Type{
	MountTypeBind:   mount.TypeBind, // Deprecated, it will be removed in a future release
	MountTypeVolume: mount.TypeVolume,
	MountTypeTmpfs:  mount.TypeTmpfs,
	MountTypePipe:   mount.TypeNamedPipe,
	MountTypeImage:  mount.TypeImage,
}

// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
// BindMounter can optionally be implemented by mount sources
// to support advanced scenarios based on mount.BindOptions
type BindMounter interface {
	GetBindOptions() *mount.BindOptions
}

// VolumeMounter can optionally be implemented by mount sources
// to support advanced scenarios based on mount.VolumeOptions
type VolumeMounter interface {
	GetVolumeOptions() *mount.VolumeOptions
}

// TmpfsMounter can optionally be implemented by mount sources
// to support advanced scenarios based on mount.TmpfsOptions
type TmpfsMounter interface {
	GetTmpfsOptions() *mount.TmpfsOptions
}

// ImageMounter can optionally be implemented by mount sources
// to support advanced scenarios based on mount.ImageOptions
type ImageMounter interface {
	ImageOptions() *mount.ImageOptions
}

// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
type DockerBindMountSource struct {
	*mount.BindOptions

	// HostPath is the path mounted into the container
	// the same host path might be mounted to multiple locations within a single container
	HostPath string
}

// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
func (s DockerBindMountSource) Source() string {
	_ = "STUB: not implemented"

	// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
	return ""
}

func (DockerBindMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// Deprecated: use Files or HostConfigModifier in the ContainerRequest, or copy files container APIs to make containers portable across Docker environments
	new(MountType)
}

func (s DockerBindMountSource) GetBindOptions() *mount.BindOptions {
	_ = "STUB: not implemented"
	return nil
}

type DockerVolumeMountSource struct {
	*mount.VolumeOptions

	// Name refers to the name of the volume to be mounted
	// the same volume might be mounted to multiple locations within a single container
	Name string
}

func (s DockerVolumeMountSource) Source() string { _ = "STUB: not implemented"; return "" }

func (DockerVolumeMountSource) Type() MountType { _ = "STUB: not implemented"; return *new(MountType) }

func (s DockerVolumeMountSource) GetVolumeOptions() *mount.VolumeOptions {
	_ = "STUB: not implemented"
	return nil
}

type DockerTmpfsMountSource struct {
	GenericTmpfsMountSource
	*mount.TmpfsOptions
}

func (s DockerTmpfsMountSource) GetTmpfsOptions() *mount.TmpfsOptions {
	_ = "STUB: not implemented"
	return nil

	// DockerImageMountSource is a mount source for an image
}

type DockerImageMountSource struct {
	// imageName is the image name
	imageName string

	// subpath is the subpath to mount the image into
	subpath string
}

// NewDockerImageMountSource creates a new DockerImageMountSource
func NewDockerImageMountSource(imageName string, subpath string) DockerImageMountSource {
	_ = "STUB: not implemented"
	return *new(DockerImageMountSource)
}

// Validate validates the source of the mount, ensuring that the subpath is a relative path
func (s DockerImageMountSource) Validate() error { _ = "STUB: not implemented"; return nil }

// ImageOptions returns the image options for the image mount
func (s DockerImageMountSource) ImageOptions() *mount.ImageOptions {
	_ = "STUB: not implemented"
	return nil
}

// Source returns the image name for the image mount
func (s DockerImageMountSource) Source() string {
	_ = "STUB: not implemented"

	// Type returns the mount type for the image mount
	return ""
}

func (s DockerImageMountSource) Type() MountType {
	_ = "STUB: not implemented"
	return *

	// PrepareMounts maps the given []ContainerMount to the corresponding
	// []mount.Mount for further processing
	new(MountType)
}

func (m ContainerMounts) PrepareMounts() []mount.Mount { _ = "STUB: not implemented"; return nil }

// mapToDockerMounts maps the given []ContainerMount to the corresponding
// []mount.Mount for further processing
func mapToDockerMounts(containerMounts ContainerMounts) []mount.Mount {
	_ = "STUB: not implemented"
	return nil
}

// The provided source type has no custom options
