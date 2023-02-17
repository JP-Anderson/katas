package wordchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsEmptyListForInvalidStartEnd(t *testing.T) {
	result := Solve("cat", "sheep")
	assert.Empty(t, result)
}

func TestDiffOfWords(t *testing.T) {
	type testCase struct {
		s1, s2 string
		expectedOut int
	}
	
	cases := []testCase{
		{
			s1: "dog", s2: "cat",
			expectedOut: 3,
		},
		{
			s1: "dog", s2: "dog",
			expectedOut: 0,
		},
		{
			s1: "dog", s2: "dig",
			expectedOut: 1,
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expectedOut, diffOfWords(tc.s1, tc.s2))
	}
}
