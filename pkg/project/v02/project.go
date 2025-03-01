package v02

import (
	"github.com/BurntSushi/toml"

	"github.com/david-caro/cnb-prepare/pkg/project/types"
)

type Buildpacks struct {
	Include []string          `toml:"include"`
	Exclude []string          `toml:"exclude"`
	Group   []types.Buildpack `toml:"group"`
	Build   Build             `toml:"build"`
	Builder string            `toml:"builder"`
}

type Build struct {
	Env []types.EnvVar `toml:"env"`
}

type Project struct {
	Name          string                 `toml:"name"`
	Licenses      []types.License        `toml:"licenses"`
	Metadata      map[string]interface{} `toml:"metadata"`
	SchemaVersion string                 `toml:"schema-version"`
}

type IO struct {
	Buildpacks Buildpacks `toml:"buildpacks"`
}

type Descriptor struct {
	Project Project `toml:"_"`
	IO      IO      `toml:"io"`
}

func NewDescriptor(projectTomlContents string) (types.Descriptor, error) {
	versionedDescriptor := &Descriptor{}
	_, err := toml.Decode(projectTomlContents, &versionedDescriptor)
	if err != nil {
		return types.Descriptor{}, err
	}

	return types.Descriptor{
		Project: types.Project{
			Name:     versionedDescriptor.Project.Name,
			Licenses: versionedDescriptor.Project.Licenses,
		},
		Build: types.Build{
			Include:    versionedDescriptor.IO.Buildpacks.Include,
			Exclude:    versionedDescriptor.IO.Buildpacks.Exclude,
			Buildpacks: versionedDescriptor.IO.Buildpacks.Group,
			Env:        versionedDescriptor.IO.Buildpacks.Build.Env,
			Builder:    versionedDescriptor.IO.Buildpacks.Builder,
		},
		Metadata:      versionedDescriptor.Project.Metadata,
		SchemaVersion: types.ParseVersion("0.2"),
	}, nil
}
