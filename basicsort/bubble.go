package basicsort

func BubbleSort(arr Element) Element {
	for i := 0; i < arr.Len() - 1; i++ {
		for j := i + 1; j < arr.Len(); j++ {
			if arr.By(arr.Slice[i], arr.Slice[j]) {
				arr.Swap(i, j)
			}
		}
	}
	return arr
}