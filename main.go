package main

import (
	"fmt"
//	"sort"
)

type people []string

/*
func (p people) String() string {
	return fmt.Sprintf("%s: %d", p)
}

type ByName []people

func (a ByName) Len() int { return len(a)}
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a ByName) Less(i, j int) bool { return a[i] < a[j]}

*/

func main() {	
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}


	fmt.Println(studyGroup)
//	sort.Sort(ByName(studyGroup))
	fmt.Println(studyGroup[1])
}
