package pkg

import "reflect"

type Scanner[T Library] interface {
	String() string
	Filename() string
	Collect(path string) (proc Processor[T], err error)
}

type Library interface {
	String() string
}

type LibrarySet[T Library] []T

func (l LibrarySet[T]) contains(library T) bool {
	for _, trusted := range l {
		if reflect.DeepEqual(library, trusted) {
			return true
		}
	}
	return false
}
