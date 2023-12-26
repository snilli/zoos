package set

type AbstractSet[T comparable] struct {
	m map[T]struct{}
}

type Set[T comparable] interface {
	Add(value T)
	Values() []T
	Contains(value T) bool
	Remove(value T)
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
	SymmetricDifference(other Set[T]) Set[T]
	IsSubsetOf(other Set[T]) bool
	IsSupersetOf(other Set[T]) bool
	Copy() Set[T]
	Length() int
}

func NewSet[T comparable](init []T) Set[T] {
	s := make(map[T]struct{})
	for _, value := range init {
		s[value] = struct{}{}
	}

	return &AbstractSet[T]{
		m: make(map[T]struct{}),
	}
}

func (s *AbstractSet[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *AbstractSet[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *AbstractSet[T]) Contains(value T) bool {
	_, c := s.m[value]
	return c
}

func (s *AbstractSet[T]) Values() []T {
	var res []T
	for key := range s.m {
		res = append(res, key)
	}
	return res
}

func (s *AbstractSet[T]) IsSubsetOf(other Set[T]) bool {
	for _, value := range s.Values() {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (s *AbstractSet[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s *AbstractSet[T]) Length() int {
	return len(s.m)
}

func (s *AbstractSet[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T](nil)
	for _, value := range s.Values() {
		result.Add(value)
	}
	for _, value := range other.Values() {
		result.Add(value)
	}
	return result
}

func (s *AbstractSet[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T](nil)
	for _, value := range s.Values() {
		if other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (s *AbstractSet[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T](nil)
	for _, value := range s.Values() {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (s *AbstractSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	diff1 := s.Difference(other)
	diff2 := other.Difference(s)
	return diff1.Union(diff2)
}

func (s *AbstractSet[T]) IsSubset(other Set[T]) bool {
	for _, value := range s.Values() {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (s *AbstractSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s *AbstractSet[T]) Copy() Set[T] {
	copyMap := make(map[T]struct{}, len(s.m))
	for key := range s.m {
		copyMap[key] = struct{}{}
	}
	return &AbstractSet[T]{m: copyMap}
}
