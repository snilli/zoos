package list

func chunk[T interface{}](slice []T, chunkSize int) [][]T {
	batches := make([][]T, 0, (len(slice)+chunkSize-1)/chunkSize)
	for chunkSize < len(slice) {
		slice, batches = slice[chunkSize:], append(batches, slice[0:chunkSize:chunkSize])
	}

	return append(batches, slice)
}

func merge[T interface{}](chunks [][]T) []T {
	list := make([]T, 0)
	for _, chunk := range chunks {
		list = append(list, chunk...)
	}

	return list
}
