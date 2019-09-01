package main

import "sortExercise/basicsort"

type Student struct {
	Id   int
	Name string
}

var Students = basicsort.Element{
	Slice: []interface{}{Student{58, "wangsen"}, Student{76, "zuyi"}, Student{49, "yuhang"}, Student{41, "pingyong"}, Student{83, "xinrui"}, Student{65, "deqing"}, Student{77, "siqing"}, Student{50, "nima"}},
	By: func(a, b interface{})bool {return a.(Student).Id > b.(Student).Id},
}