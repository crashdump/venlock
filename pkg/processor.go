package pkg

import (
	"reflect"
)

// TODO: Use go Generic's `comparable` instead of reflection
type Processor[T any] struct {
	Inventory LibrarySet[T]
}

func (v *Processor[T]) List() LibrarySet[T] {
	return v.Inventory
}

func (v *Processor[T]) Validate(catalogue LibrarySet[T]) (foreign LibrarySet[T], valid bool) {
	for _, invLib := range v.Inventory {
		for _, catLib := range catalogue {
			if reflect.ValueOf(invLib).Interface() == reflect.ValueOf(catLib).Interface() {
				continue
			}
			//if invLib == catLib {
			//	continue
			//}
			foreign = append(foreign, invLib)
		}
	}
	return foreign, true
}
