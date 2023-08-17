package npm

import (
	"os"

	"github.com/crashdump/venlock/pkg"
)

type Npm[T Library] struct{}

var _ pkg.Scanner[Library] = Npm[Library]{}

func (Npm[T]) String() string {
	return "npm"
}

func (Npm[T]) Filename() string {
	return "package.json"
}

func (Npm[T]) Collect(path string) (proc pkg.Processor[Library], err error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return proc, err
	}

	pjson, err := parse(bytes)
	if err != nil {
		return proc, err
	}

	for name := range pjson.Dependencies {
		proc.Found = append(proc.Found, Library{
			Name: name,
		})
	}

	return proc, nil
}
