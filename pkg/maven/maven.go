package maven

import (
	"github.com/vifraa/gopom"

	"github.com/crashdump/libguardian/pkg"
)

type Maven[T Library] struct{}

type Library struct {
	GroupId    string `json:"group_id"`
	ArtefactId string `json:"artefact_id"`
}

var _ pkg.PackageManager[Library] = Maven[Library]{}

func (m Maven[T]) Name() string {
	return "maven"
}

func (m Maven[T]) Filename() string {
	return "pom.xml"
}

func (m Maven[T]) Collect(path string) (proc pkg.Processor[T], err error) {
	pom, err := gopom.Parse(path)
	if err != nil {
		return proc, err
	}

	for _, dep := range pom.Dependencies {
		proc.Inventory = append(proc.Inventory, T{
			GroupId:    dep.GroupID,
			ArtefactId: dep.ArtifactID,
		})
	}

	return proc, nil
}
