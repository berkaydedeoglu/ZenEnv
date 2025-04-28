package zenvar

import "fmt"

type ZenVar struct {
	Name     string
	Value    string
	Project  string
	Platform string
	Tag      string
}

func NewZenVar(
	name string,
	value string,
	project string,
	platform string,
	tag string,
) *ZenVar {
	zv := createDefaultZenvar()

	zv.Name = name
	zv.Value = value
	zv.Project = project
	zv.Platform = platform
	zv.Tag = tag

	return zv
}

func createDefaultZenvar() *ZenVar {
	return &ZenVar{
		Project:  DefaultProject,
		Platform: DefaultPlatform,
		Tag:      DefaultTag,
	}
}

func (z *ZenVar) GenerateKey() string {
	return fmt.Sprintf("%s.%s.%s.%s",
		z.Project,
		z.Platform,
		z.Name,
		z.Tag)
}
