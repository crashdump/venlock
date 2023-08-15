package pkg

import (
	"errors"
)

type LibGuardian[T any] struct {
	SourcePath string
	Scanner    PackageManager[T]
	Config     Config[T]
}

func NewLibGuardian[T any](sourcePath string, scanner PackageManager[T]) LibGuardian[T] {
	return LibGuardian[T]{
		SourcePath: sourcePath,
		Scanner:    scanner,
	}
}

// Enumerate scans for known manifests in the target directory. It will then list all the
// libraries found in this directory.
func (l *LibGuardian[T]) Enumerate() (libraries LibrarySet[T], err error) {

	manifests, err := FindManifests(l.Scanner.Filename(), l.SourcePath)
	if err != nil {
		return libraries, err
	}

	for _, manifest := range manifests {
		proc, err := l.Scanner.Collect(manifest)
		if err != nil {
			return libraries, err
		}

		libraries = append(libraries, proc.List()...)
	}

	return libraries, nil
}

// Generate scans for known manifests in the target directory. It will then generate a config
// file based on the libraries found in this directory.
func (l *LibGuardian[T]) Generate(configPath string) (foreign LibrarySet[T], err error) {
	manifests, err := FindManifests(l.Scanner.Filename(), l.SourcePath)
	if err != nil {
		return foreign, err
	}

	for _, manifest := range manifests {
		_, err := l.Scanner.Collect(manifest)
		if err != nil {
			return foreign, err
		}

		gaps, err := l.Scanner.Collect(l.SourcePath)
		if err != nil {
			return foreign, err
		}

		foreign = append(foreign, gaps.Inventory...)
	}

	err = l.Config.Save(configPath)
	if err != nil {
		return foreign, err
	}

	return foreign, nil
}

// Enforce scans for known manifests in the target directory. If an unknown library is found,
// the program will exit with error code 1.
func (l *LibGuardian[T]) Enforce() (foreign LibrarySet[T], err error) {
	foreign = make(LibrarySet[T], 0)

	manifests, err := FindManifests(l.Scanner.Filename(), l.SourcePath)
	if err != nil {
		return foreign, err
	}

	for _, manifest := range manifests {
		proc, err := l.Scanner.Collect(manifest)
		if err != nil {
			return foreign, err
		}

		gaps, valid := proc.Validate(l.Config.Catalogue[l.Scanner.Name()])
		if !valid {
			return foreign, errors.New("found unexpected libraries")
		}

		foreign = append(foreign, gaps...)
	}

	return foreign, nil
}
