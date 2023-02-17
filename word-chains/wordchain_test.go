package wordchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsEmptyListForInvalidStartEnd(t *testing.T) {
	result := Solve("cat", "sheep")
	assert.Empty(t, result)
}

func TestSolveCatToCot(t *testing.T) {
	result := Solve("cat", "cot")
	assert.Len(t, result, 2) 	
	assert.Equal(t, result[0], "cat")
 	assert.Equal(t, result[1], "cot")
}

func TestSolveBatToBum(t *testing.T) {
	result := Solve("bat", "bum")
	assert.Len(t, result, 3)
	assert.Equal(t, result[0], "bat")
	assert.Equal(t, 1, diffOfWords("bat", result[1]))
	assert.Equal(t, 1, diffOfWords("bum", result[1]))
	assert.Equal(t, result[2], "bum")
}

func TestSolve5LetterWords(t *testing.T) {
	result := Solve("carve", "halve")
	assert.Len(t, result, 3)

	result = Solve("prose", "prank")
	assert.Len(t, result, 3)
}

func TestSolveCatToCog(t *testing.T) {
	result := Solve("cat", "cog")
	assert.Len(t, result, 3)
	assert.Equal(t, result[0], "cat")
	assert.Equal(t, 1, diffOfWords("cat", result[1]))
	assert.Equal(t, 1, diffOfWords("cog", result[1]))
	assert.Equal(t, result[2], "cog")
}

func TestSolveCatToDog(t *testing.T) {
	t.Skip()
	result := Solve("cat", "dog")
	assert.Len(t, result, 4)	
	assert.Equal(t, result[0], "cat")
	assert.Equal(t, 1, diffOfWords("cat", result[1]))
	assert.Equal(t, 1, diffOfWords("dog", result[2]))
	assert.Equal(t, result[3], "dog")
}

func TestVisitRow(t *testing.T) {
	rows := [][]string{
		{ "cat" },
		{ "ignore", "multivalues" },
		{ "dog" },
		{ "fish" },
		{ "bush" },
		{ "apps" },
		{ "and" },
		{ "wolf" },
		{ "lock" },
		{ "jock" },
		{ "leet" },
		{ "leot" },
		{ "leok" },
		{ "look" },
	}
	targetWord := "look"
	startWord := "meet"

	visitor := newVisitor(startWord, targetWord)
	assert.Equal(t, len(targetWord), visitor.wordLength)
	for _, row := range rows {
		visitor.VisitRow(row)
	}
	assert.Contains(t, visitor.diffCharCountToWords, 1)
	assert.Contains(t, visitor.diffCharCountToWords, 2)
	assert.Contains(t, visitor.diffCharCountToWords, 3)
	assert.Contains(t, visitor.diffCharCountToWords, 4)	
		
	assert.Equal(t, []string{ "leet" }, visitor.diffCharCountToWords[1])
	assert.Equal(t, []string{ "leot" }, visitor.diffCharCountToWords[2])
	assert.Equal(t, []string{ "leok" }, visitor.diffCharCountToWords[3])
	assert.Equal(t, []string{ "wolf", "lock", "jock" }, visitor.diffCharCountToWords[4])
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
