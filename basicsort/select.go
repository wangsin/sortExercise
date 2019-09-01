package basicsort

func SelectSort(arr Element) Element {
	for i := 0; i < arr.Len(); i++ {
		min := arr.Slice[i]
		for j := i + 1; j < arr.Len(); j++ {
			if arr.By(min, arr.Slice[j]) {
				min, arr.Slice[j] = arr.Slice[j], min
			}
		}
		arr.Slice[i] = min
	}

	return arr
}