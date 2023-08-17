package pkg

type Venlock[T Library] struct {
	SourcePath string
	Scanner    Scanner[T]
}

func NewVenlock[T Library](sourcePath string, scanner Scanner[T]) Venlock[T] {
	return Venlock[T]{
		SourcePath: sourcePath,
		Scanner:    scanner,
	}
}

// Enumerate scans for known manifests in the target directory. It will then list all the
// libraries found in this directory.
func (l *Venlock[T]) Enumerate() (libraries LibrarySet[T], err error) {
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
func (l *Venlock[T]) Generate(configPath string) (err error) {
	manifests, err := FindManifests(l.Scanner.Filename(), l.SourcePath)
	if err != nil {
		return err
	}

	for _, manifest := range manifests {
		_, err := l.Scanner.Collect(manifest)
		if err != nil {
			return err
		}

		proc, err := l.Scanner.Collect(l.SourcePath)
		if err != nil {
			return err
		}

		err = proc.Save(configPath, l.Scanner.String())
		if err != nil {
			return err
		}
	}

	return nil
}

// Enforce scans for known manifests in the target directory. If an unknown library is found,
// the program will exit with error code 1.
func (l *Venlock[T]) Enforce(catalogue LibrarySet[T]) (foreign LibrarySet[T], err error) {
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

		found, _ := proc.Validate(catalogue)
		foreign = append(foreign, found...)
	}

	return foreign, nil
}
