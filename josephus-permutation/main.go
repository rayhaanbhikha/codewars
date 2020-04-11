package main

func main() {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Josephus(items, 2)
}

func Josephus(items []interface{}, k int) []interface{} {
	newItems := make([]interface{}, 0)
	for i := k - 1; len(items) != 0; i += k - 1 {
		if i >= len(items) {
			i = i % len(items)
		}
		newItems = append(newItems, items[i])
		if i+1 < len(items) {
			items = append(items[:i], items[i+1:]...)
		} else {
			items = items[:i]
		}
	}
	return newItems
}
