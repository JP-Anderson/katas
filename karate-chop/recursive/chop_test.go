package main

import (
	"fmt"
	"testing"
)

func TestChop(t *testing.T) {
	cases := []struct{
		searchItem int
		sortedSlice []int
		desiredOutput int
	}{
		{1, []int{}, -1},
		{1, []int{1}, 0},
		{1, []int{1, 2}, 0},
		{1, []int{0, 1}, 1},
		{1, []int{1, 2, 3}, 0},
		{1, []int{0, 1, 2, 3}, 1},
		{1, []int{-1, 0, 1, 2, 3}, 2},
		{1, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 6}, 
		{1, []int{-5, -4, -3, -2, -1, 0, 1}, 6}, 
		{1, []int{-5, -4, -3, -2, -1, 0, 1, 2}, 6}, 
		{1, []int{-3, -2, -1, 0, 1, 2, 3}, 4}, 
	}

	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%d, %#v", testCase.searchItem, testCase.sortedSlice), func(t *testing.T) {
			result := Chop(testCase.searchItem, testCase.sortedSlice)
			if testCase.desiredOutput != result {
				t.Errorf("got %d, want %d", result, testCase.desiredOutput)
			}
		})
	}
}
