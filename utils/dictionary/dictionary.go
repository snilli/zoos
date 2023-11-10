package dictionary

import "fmt"

type AbstractDictionary[K comparable, T interface{}] struct {
	m map[K]T
}

type Dictionary[K comparable, T interface{}] interface {
	Get(key K) (T, bool)
	Set(key K, value T)
	Contains(key K) bool
	Remove(key K)
	Keys() *[]K
	Values() *[]T
	Entries() map[K]T
	Copy() Dictionary[K, T]
}

func NewDictionary[K comparable, T interface{}]() Dictionary[K, T] {
	return &AbstractDictionary[K, T]{
		m: make(map[K]T),
	}
}

func (d *AbstractDictionary[K, T]) Get(key K) (T, bool) {
	val, ok := d.m[key]
	return val, ok
}

func (d *AbstractDictionary[K, T]) Set(key K, value T) {
	d.m[key] = value
}

func (d *AbstractDictionary[K, T]) Contains(key K) bool {
	_, c := d.m[key]
	return c
}

func (d *AbstractDictionary[K, T]) Remove(key K) {
	delete(d.m, key)
}

func (d *AbstractDictionary[K, T]) Keys() *[]K {
	var keys []K
	for key := range d.m {
		keys = append(keys, key)
	}
	return &keys
}

func (d *AbstractDictionary[K, T]) Values() *[]T {
	var values []T
	for _, value := range d.m {
		values = append(values, value)
	}
	return &values
}

func (d *AbstractDictionary[K, T]) Entries() map[K]T {
	return d.m
}

func (d *AbstractDictionary[K, T]) Copy() Dictionary[K, T] {
	copyMap := make(map[K]T, len(d.m))
	for k, v := range d.m {
		copyMap[k] = v
	}
	return &AbstractDictionary[K, T]{m: copyMap}
}

type B struct {
	V string
	B string
}
type A struct {
	V Dictionary[string, B]
}

func v() {
	a := A{
		V: NewDictionary[string, B](),
	}
	a.V.Set("kook", B{"a", "a"})
	a.V.Set("kook1", B{"a", "a"})
	a.V.Set("kook2", B{"a", "a"})
	a.V.Set("kook3", B{"a", "a"})
	a.V.Set("kook4", B{"a", "a"})
	a.V.Set("kook5", B{"a", "a"})
	for c, b := range *a.V.Keys() {
		fmt.Println(c, b)
	}
}
