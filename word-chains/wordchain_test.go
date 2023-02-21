package wordchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsEmptyListForInvalidStartEnd(t *testing.T) {
	result := Solve("cat", "sheep")
	assert.Empty(t, result)

	result = Solve("harries", "markets")
	assert.Empty(t, result)
}

func TestSameWordChainBothWays(t *testing.T) {
	result := Solve("gold", "lead")
	result2 := Solve("lead", "gold")
	for i := 0; i < 4; i++ {
		assert.Equal(t, result[i], result2[3-i])
	}
}

func TestSolveRubyToCode(t *testing.T) {
	result := Solve("ruby", "code")
	assert.Equal(t, []string{ "ruby", "rubs", "robs", "rods", "rode", "code" },  result)
}

func TestSolveCatToCot(t *testing.T) {
	result := Solve("cat", "cot")
	assert.Len(t, result, 2) 	
	assert.Equal(t, result[0], "cat")
 	assert.Equal(t, result[1], "cot")
}

func TestSolveBatToBum(t *testing.T) {
	result := Solve("bat", "bum")
	assert.Equal(t, []string{ "bat", "bam", "bum" }, result)
}

func TestSolveChainsOfLength(t *testing.T) {
	// 1 Diff
	result := Solve("chase", "chast")
	assert.Equal(t, []string{ "chase", "chast" }, result)		
	
	// 2 Diffs
	result = Solve("carve", "halve")
	assert.Equal(t, []string{"carve", "calve", "halve"}, result)

	// 3 Diffs
	result = Solve("crank", "doink")
	assert.Equal(t, []string{"crank", "drank", "drink", "doink"}, result)

	// 4 Diffs
	result = Solve("curve", "salvo")
	assert.Equal(t, []string{"curve", "carve", "calve", "salve", "salvo"}, result)
	
	// 5 Diffs
	result = Solve("table", "corps")
	assert.Equal(t, []string{"table", "cable", "carle", "carls", "carps", "corps"}, result)
	
	result = Solve("atlases", "unlaced")
	assert.Equal(t, []string{"atlases", "anlases", "anlaces", "unlaces", "unlaced"}, result)
	
}

func TestSolveCatToCog(t *testing.T) {
	result := Solve("cat", "cog")
	assert.Len(t, result, 3)
	assert.Equal(t, result[0], "cat")
	assert.Equal(t, 1, diffOfWords("cat", result[1]))
	assert.Equal(t, 1, diffOfWords("cog", result[1]))
	assert.Equal(t, []string{ "cat", "cot", "cog" }, result)
}

func TestSolveCatToDog(t *testing.T) {
	result := Solve("cat", "dog")
	assert.Equal(t, []string{ "cat", "cot", "cog", "dog" }, result)
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
