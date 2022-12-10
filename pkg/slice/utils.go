package slice

import "github.com/DiasOrazbaev/kazanexpress-test-task/internal/batch"

func ChunkSlice(slice []batch.Item, chunkSize int) [][]batch.Item {
	var chunks [][]batch.Item
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
