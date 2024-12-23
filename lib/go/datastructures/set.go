package datastructures

type Set[T comparable] struct {
	Items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		Items: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(item T) {
	s.Items[item] = struct{}{}
}

func (s *Set[T]) AddM(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

func (s *Set[T]) Remove(item T) {
	delete(s.Items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.Items[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.Items)
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.Items))
	for v := range s.Items {
		slice = append(slice, v)
	}
	return slice
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := NewSet[T]()
	for v := range s.Items {
		union.Add(v)
	}
	for v := range other.Items {
		union.Add(v)
	}
	return union
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := NewSet[T]()
	for item := range s.Items {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}
