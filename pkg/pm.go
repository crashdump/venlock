package pkg

type PackageManager[T any] interface {
	Name() string
	Filename() string
	Collect(path string) (proc Processor[T], err error)
}

type LibrarySet[T any] []T
