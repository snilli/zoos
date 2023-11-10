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

func Map[T interface{}, E interface{}](data []T, f func(T) E) []E {
	mapped := make([]E, len(data))
	for i, e := range data {
		mapped[i] = f(e)
	}

	return mapped
}

func Filter[T interface{}](data []T, f func(T) bool) []T {
	fltd := make([]T, 0, len(data))
	for _, e := range data {
		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}
