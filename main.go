package main

import (
	"fmt"
	"sortExercise/basicsort"
)

func main() {
	fmt.Println(Students.Slice)
	//basicsort.BubbleSort(Students)
	//basicsort.InsertSort(Students)
	basicsort.ShellSort(Students)
	//basicsort.SelectSort(Students)
	//basicsort.QuickSort(Students, 0, Students.Len() - 1)
	fmt.Println(Students.Slice)
}
