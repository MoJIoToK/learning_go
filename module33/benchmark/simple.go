package benchmark

// Simple возвращает номер элемента в массиве или -1.
// Используется простой поиск.
func Simple(data []int, item int) int {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}
