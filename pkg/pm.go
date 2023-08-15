package pkg

import "reflect"

type PackageManager[T any] interface {
	Name() string
	Filename() string
	Collect(path string) (proc Processor[T], err error)
}

type LibrarySet[T any] []T

func (l LibrarySet[T]) contains(library T) bool {
	for _, trusted := range l {
		if reflect.DeepEqual(library, trusted) {
			return true
		}
	}
	return false
}
