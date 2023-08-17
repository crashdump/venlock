package gomod

import (
	"os"

	"golang.org/x/mod/modfile"

	"github.com/crashdump/venlock/pkg"
)

type GoMod[T Library] struct{}

var _ pkg.Scanner[Library] = GoMod[Library]{}

func (GoMod[T]) String() string {
	return "gomod"
}

func (GoMod[T]) Filename() string {
	return "go.mod"
}

func (GoMod[T]) Collect(path string) (proc pkg.Processor[Library], err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return proc, err
	}

	gomod, err := modfile.ParseLax(path, data, nil)
	if err != nil {
		return proc, err
	}

	for _, dep := range gomod.Require {
		proc.Found = append(proc.Found, Library{
			Module: dep.Mod.Path,
		})
	}

	return proc, nil
}
