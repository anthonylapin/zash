package util

type Set[T comparable] map[T]struct{}

// Add inserts an element into the Set
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (set Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(set))

	for key := range set {
		keys = append(keys, key)
	}

	return keys
}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T])
	
	for _, v := range values {
		set.Add(v)
	}
	
	return set
}