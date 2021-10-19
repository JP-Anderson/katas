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
