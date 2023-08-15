package gomod

import (
	"os"

	"golang.org/x/mod/modfile"

	"github.com/crashdump/libguardian/pkg"
)

type GoMod[T Library] struct{}

type Library struct {
	Module string `json:"module"`
}

var _ pkg.PackageManager[Library] = GoMod[Library]{}

func (g GoMod[T]) Name() string {
	return "gomod"
}

func (g GoMod[T]) Filename() string {
	return "go.mod"
}

func (g GoMod[T]) Collect(path string) (proc pkg.Processor[T], err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return proc, err
	}

	gomod, err := modfile.ParseLax(path, data, nil)
	if err != nil {
		return proc, err
	}

	for _, dep := range gomod.Require {
		proc.Inventory = append(proc.Inventory, T{
			Module: dep.Mod.Path,
		})
	}

	return proc, nil
}
