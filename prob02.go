package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.StringSlice(s))
	fmt.Printf("Regular Set: %s\n",s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Printf("Reverse Set: %s\n",s)
}