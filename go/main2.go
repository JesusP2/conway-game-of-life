package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	removeIndex := RemoveIndex(all, 5)

	printSlice(all)
	printSlice(removeIndex)

	removeIndex[8] = 999
	printSlice(all)
	printSlice(removeIndex)

	idk := make([]int, 0)
	printSlice(idk)
	idk2 := append(idk, 1, 2, 3, 4, 5)
	idk3 := append(idk2, 6)
	idk3[0] = 999
	idk4 := append(idk2, 6, 7)
	idk3[1] = 999

	printSlice(idk)
	printSlice(idk2)
	printSlice(idk3)
	printSlice(idk4)

	// array := [5]int{1, 2, 3}
	// s1 := array[:1]
	// s2 := append(s1, 1, 2, 3, 4, 5)
	// fmt.Println(array) // [1, 2, 3]
	// printSlice(s1)     // [1]
	// printSlice(s2)     // [1, 1, 2]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %p, %v\n", len(s), cap(s), s, s)
}
