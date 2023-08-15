package npm

import (
	"os"

	"github.com/crashdump/libguardian/pkg"
)

type Npm[T Library] struct{}

type Library struct {
	Name string `json:"name"`
}

var _ pkg.PackageManager[Library] = Npm[Library]{}

func (n Npm[T]) Name() string {
	return "npm"
}

func (n Npm[T]) Filename() string {
	return "package.json"
}

func (n Npm[T]) Collect(path string) (proc pkg.Processor[T], err error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return proc, err
	}

	pjson, err := parse(bytes)
	if err != nil {
		return proc, err
	}

	for name := range pjson.Dependencies {
		proc.Inventory = append(proc.Inventory, T{
			Name: name,
		})
	}

	return proc, nil
}
