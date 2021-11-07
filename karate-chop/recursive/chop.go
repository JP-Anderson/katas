package main

import (
	"fmt"
)

func Chop(searchItem int, sortedIntSlice []int) int {
	if len(sortedIntSlice) == 0 {
		return -1
	}
	middleIndex := len(sortedIntSlice)/2
	middleValue := sortedIntSlice[middleIndex]
	fmt.Printf("sorted int slice = %#v\n middle index = %d\n mid val = %d\n", sortedIntSlice, middleIndex, middleValue)
	if searchItem == middleValue {
		return middleIndex
	} else if searchItem < middleValue {
		return Chop(searchItem, sortedIntSlice[0:middleIndex])
	} else if searchItem > middleValue {
		return middleIndex + Chop(searchItem, sortedIntSlice[middleIndex: len(sortedIntSlice)])
	}
	return middleIndex
}

