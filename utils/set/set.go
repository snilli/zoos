package set

type set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *set[T] {
	return &set[T]{
		m: make(map[T]struct{}),
	}
}

func NewSetWithValues[T comparable](init *[]T) *set[T] {
	s := NewSet[T]()
	if len(*init) > 0 {
		for _, value := range *init {
			s.Add(value)
		}
	}
	return s
}

func (s *set[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *set[T]) Contains(value T) bool {
	_, c := s.m[value]
	return c
}

func (s *set[T]) Get() []T {
	var res []T
	for key := range s.m {
		res = append(res, key)
	}
	return res
}

func (s *set[T]) Union(other *set[T]) *set[T] {
	result := NewSet[T]()
	for _, value := range s.Get() {
		result.Add(value)
	}
	for _, value := range other.Get() {
		result.Add(value)
	}
	return result
}

func (s *set[T]) Intersection(other *set[T]) *set[T] {
	result := NewSet[T]()
	for _, value := range s.Get() {
		if other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (s *set[T]) Difference(other *set[T]) *set[T] {
	result := NewSet[T]()
	for _, value := range s.Get() {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (s *set[T]) SymmetricDifference(other *set[T]) *set[T] {
	result := NewSet[T]()
	for _, value := range s.Get() {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	for _, value := range other.Get() {
		if !s.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (s *set[T]) IsSubsetOf(other *set[T]) bool {
	for _, value := range s.Get() {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (s *set[T]) IsSupersetOf(other *set[T]) bool {
	return other.IsSubsetOf(s)
}
