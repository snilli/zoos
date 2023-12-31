package list

func Chunk[T interface{}](slice []T, chunkSize int) [][]T {
	batches := make([][]T, 0, (len(slice)+chunkSize-1)/chunkSize)
	for chunkSize < len(slice) {
		slice, batches = slice[chunkSize:], append(batches, slice[0:chunkSize:chunkSize])
	}

	return append(batches, slice)
}

func Merge[T interface{}](chunks [][]T) []T {
	list := make([]T, 0)
	for _, chunk := range chunks {
		list = append(list, chunk...)
	}

	return list
}

func Map[T, E interface{}](data []T, f func(slice T, i int) E) []E {
	mapped := make([]E, len(data))
	for i, e := range data {
		mapped[i] = f(e, i)
	}

	return mapped
}

func Filter[T interface{}](data []T, f func(slice T, i int) bool) []T {
	fltd := make([]T, 0, len(data))
	for i, e := range data {
		if f(e, i) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}

func Find[T interface{}](data []T, f func(slice T, i int) bool) bool {
	for i, e := range data {
		if f(e, i) {
			return true
		}
	}

	return false
}

func IndexOf[T interface{}](data []T, f func(slice T, i int) bool) int {
	for i, e := range data {
		if f(e, i) {
			return i
		}
	}

	return -1
}

func Reduce[T, E interface{}](slice []T, initial E, combiner func(E, T, int) E) E {
	result := initial
	for index, element := range slice {
		result = combiner(result, element, index)
	}
	return result
}
