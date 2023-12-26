package dictionary

import (
	"sort"
)

type AbstractDictionary[K comparable, T interface{}] struct {
	m map[K]T
}

type DictionaryEntry[K comparable, T interface{}] struct {
	K K
	V T
}

type Dictionary[K comparable, T interface{}] interface {
	Get(key K) (T, bool)
	Set(key K, value T)
	Contains(key K) bool
	Remove(key K)
	Keys() []K
	Values() []T
	Entries() []DictionaryEntry[K, T]
	Copy() Dictionary[K, T]
	Length() int
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

func (d *AbstractDictionary[K, T]) Length() int {
	return len(d.m)
}

func (d *AbstractDictionary[K, T]) Contains(key K) bool {
	_, c := d.m[key]
	return c
}

func (d *AbstractDictionary[K, T]) Remove(key K) {
	delete(d.m, key)
}

func (d *AbstractDictionary[K, T]) Keys() []K {
	var keys []K
	entries := d.sort()
	for _, entry := range entries {
		keys = append(keys, entry.K)
	}
	return keys
}

func (d *AbstractDictionary[K, T]) Values() []T {
	var values []T
	entries := d.sort()
	for _, entry := range entries {
		values = append(values, entry.V)
	}
	return values
}

func (d *AbstractDictionary[K, T]) Entries() []DictionaryEntry[K, T] {
	return d.sort()
}

func (d *AbstractDictionary[K, T]) Copy() Dictionary[K, T] {
	copyMap := make(map[K]T, len(d.m))
	for k, v := range d.m {
		copyMap[k] = v
	}

	return &AbstractDictionary[K, T]{m: copyMap}
}

func (d *AbstractDictionary[K, T]) sort() []DictionaryEntry[K, T] {
	var keys []interface{}
	for key := range d.m {
		keys = append(keys, key)
	}

	d.sortKeys(keys)

	res := []DictionaryEntry[K, T]{}

	for _, key := range keys {
		value := d.m[key.(K)]
		res = append(res, DictionaryEntry[K, T]{K: key.(K), V: value})
	}

	return res
}

func (d *AbstractDictionary[K, T]) sortKeys(keys []interface{}) {
	sort.Slice(keys, func(i, j int) bool {
		switch keys[i].(type) {
		case int:
			return keys[i].(int) < keys[j].(int)
		case string:
			return keys[i].(string) < keys[j].(string)
		default:
			// Handle other types as needed
			return false
		}
	})
}
