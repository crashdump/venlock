package maven

import (
	"github.com/vifraa/gopom"

	"github.com/crashdump/venlock/pkg"
)

type Maven[T Library] struct{}

var _ pkg.Scanner[Library] = Maven[Library]{}

func (Maven[T]) String() string {
	return "maven"
}

func (Maven[T]) Filename() string {
	return "pom.xml"
}

func (Maven[T]) Collect(path string) (proc pkg.Processor[Library], err error) {
	pom, err := gopom.Parse(path)
	if err != nil {
		return proc, err
	}

	for _, dep := range pom.Dependencies {
		proc.Found = append(proc.Found, Library{
			GroupId:    dep.GroupID,
			ArtefactId: dep.ArtifactID,
		})
	}

	return proc, nil
}
