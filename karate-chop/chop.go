package main

import (
	"fmt"
)

func Chop(searchItem int, sortedIntSlice []int) int {
	if len(sortedIntSlice) == 0 {
		return -1
	}
	fmt.Printf("sorted int slice = %#v\n", sortedIntSlice)
	middleIndex := len(sortedIntSlice)/2
	middleValue := sortedIntSlice[middleIndex]
	if searchItem == middleValue {
		return middleIndex
	} else if searchItem < middleValue {
		return Chop(searchItem, sortedIntSlice[0:middleIndex])
	}
	return middleIndex
}

