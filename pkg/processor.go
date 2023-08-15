package pkg

// TODO: Use go Generic's `comparable` instead of reflection
type Processor[T any] struct {
	Found LibrarySet[T]
}

func (v *Processor[T]) List() LibrarySet[T] {
	return v.Found
}

func (v *Processor[T]) Validate(catalogue LibrarySet[T]) (foreign LibrarySet[T], compliant bool) {
	for _, found := range v.Found {
		if catalogue.contains(found) {
			continue
		}
		foreign = append(foreign, found)
	}
	if len(foreign) > 0 {
		return foreign, false
	}
	return foreign, true
}

func (v *Processor[T]) Save(path string, packageManager string) (err error) {
	var cfg Config[T]
	cfg.Catalogue[packageManager] = v.Found
	return cfg.Save(path)

}
