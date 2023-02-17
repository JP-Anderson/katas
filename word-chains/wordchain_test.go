package wordchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsEmptyListForInvalidStartEnd(t *testing.T) {
	result := Solve("cat", "sheep")
	assert.Empty(t, result)
}

func TestVisitRow(t *testing.T) {
	rows := [][]string{
		{ "cat" },
		{ "ignore", "multivalues" },
		{ "dog" },
		{ "fish" },
		{ "and" },
		{ "wolf" },
		{ "lock" },
		{ "jock" },
	}
	targetWord := "look"

	visitor := newVisitor(targetWord)
	assert.Equal(t, len(targetWord), visitor.wordLength)
	for _, row := range rows {
		visitor.VisitRow(row)
	}
	assert.Contains(t, visitor.diffCharCountToWords, 1)
	assert.Contains(t, visitor.diffCharCountToWords, 2)
	assert.Contains(t, visitor.diffCharCountToWords, 3)
	assert.Contains(t, visitor.diffCharCountToWords, 4)	
	
	assert.Contains(t, visitor.diffCharCountToWords[1], "lock")
	assert.Contains(t, visitor.diffCharCountToWords[2], "jock")
	assert.Contains(t, visitor.diffCharCountToWords[3], "wolf")
	assert.Contains(t, visitor.diffCharCountToWords[4], "fish")
	assert.Len(t, visitor.diffCharCountToWords[1], 1)
	assert.Len(t, visitor.diffCharCountToWords[2], 1)
	assert.Len(t, visitor.diffCharCountToWords[3], 1)
	assert.Len(t, visitor.diffCharCountToWords[4], 1)
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
