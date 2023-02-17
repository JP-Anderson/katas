package anagrams

import (
	"katas/csv"
)

type anagramKey struct {
	letters [26]int
}

func wordToAnagramKey(word string) (anagramKey, bool) {
	lc := [26]int{}
	for _, k := range word {
		if k < 97 || k > 122 {
			return anagramKey{}, false
		} 
		lc[k-97] = lc[k-97]+1	
	}
	return anagramKey{ letters: lc }, true
}

type AnagramRowVisitor struct {
	anagrams map[anagramKey][]string
}

func (a *AnagramRowVisitor) VisitRow(row []string) {
	key, ok := wordToAnagramKey(row[0])
	if !ok {
		return
	}
	anagrams, ok := a.anagrams[key]
	if ok {
		a.anagrams[key] = append(anagrams, row[0])
	} else {
		a.anagrams[key] = []string{ row[0] }
	}
}

func (a *AnagramRowVisitor) Groups() [][]string {
	groups := make([][]string, len(a.anagrams))
	i := 0
	for _, v := range a.anagrams {
		groups[i] = v
		i++
	}
	return groups
}

func AnagramsInFile(filename string) [][]string {
	visitor := &AnagramRowVisitor{
		anagrams: map[anagramKey][]string{},
	}
	reader := csv.NewReader(visitor)
	reader.VisitRows(filename)
	return visitor.Groups()
}
