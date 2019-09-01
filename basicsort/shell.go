package basicsort

import "fmt"

func ShellSort(arr Element) Element {
	step := arr.Len() / 2

	for step > 0 {
		for i := step; i < arr.Len(); i++ {
			for j := i; j >= step; j = j - step {
				if arr.By(arr.Slice[j - step], arr.Slice[j]) {
					arr.Swap(j - step, j)
				}
			}
			fmt.Println(arr.Slice)
		}
		step = step / 2
	}

	return arr
}
