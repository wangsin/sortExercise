package basicsort

func QuickSort(arr Element, left, right int) {
	if left < right {
		i, j, base := left, right, arr.Slice[left]
		for i < j {
			for i < j && arr.By(arr.Slice[j], base) {
				j--
			}
			arr.Slice[i], arr.Slice[j] = arr.Slice[j], arr.Slice[i]

			for i < j && arr.By(base, arr.Slice[i]) {
				i++
			}
			arr.Slice[i], arr.Slice[j] = arr.Slice[j], arr.Slice[i]
		}

		QuickSort(arr, left, i - 1)
		QuickSort(arr, i + 1, right)
	}
}
