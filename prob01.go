package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Printf("Regular Set: %s\n",studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Printf("Reverse Set: %s\n",studyGroup)
}