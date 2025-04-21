package zenvar

import "fmt"

type ZenVar struct {
	Name     string
	Value    string
	Project  string
	Platform string
	Version  string
}

func NewZenVar(
	name string,
	value string,
	project string,
	platform string,
	version string,
) *ZenVar {
	zv := createDefaultZenvar()

	zv.Name = name
	zv.Value = value
	zv.Project = project
	zv.Platform = platform
	zv.Version = version

	return zv
}

func createDefaultZenvar() *ZenVar {
	return &ZenVar{
		Project:  DefaultProject,
		Platform: DefaultPlatform,
		Version:  DefaultVersion,
	}
}

func (z *ZenVar) GenerateKey() string {
	return fmt.Sprintf("%s.%s.%s.%s", z.Project, z.Platform, z.Name, z.Version)
}
