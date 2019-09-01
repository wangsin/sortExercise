package basicsort

func InsertSort(arr Element) Element {
	for i := 0; i < arr.Len(); i++ {
		for j := i; j > 0; j-- {
			if arr.By(arr.Slice[j - 1], arr.Slice[j]) {
				arr.Swap(j - 1, j)
			}
		}
	}
	return arr
}
