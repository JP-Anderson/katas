package wordchain

import (
	"katas/csv"	
)

func Solve(start, end string) []string {
	if len(start) != len(end) {
		return []string{}
	}
	diff := diffOfWords(start, end)
	if diff == 1 {
		return []string{ start, end }
	}
	v := newVisitor(start, end)
	r := csv.NewReader(v)
	r.VisitRows("../anagrams/wordlist.txt")
	chain := make([]string, diff+1)
	chain[0] = start
	chain[len(chain)-1] = end
	return tryNextChain(chain, start, end, 0, v.diffCharCountToWords)
}

func tryNextChain(currentChain []string, start, end string, currentIx int, diffToWordsMap map[int][]string) []string {
	nextWords := diffToWordsMap[currentIx+1]
	diff := len(currentChain)-1
	if currentIx == len(currentChain)-2 {
		return currentChain
	}
	for _, w := range nextWords {
		if diffOfWords(w, end) < diff-currentIx {
			currentIx++
			currentChain[currentIx] = w
			res := tryNextChain(currentChain, start, end, currentIx, diffToWordsMap)
			if res != nil {
				return res
			}
		}
	}
	return nil
}

type wordChainSolverVisitor struct {
	wordLength int
	start, end string
	wordDiff int
	diffCharCountToWords map[int][]string
}

func newVisitor(startWord, endWord string) *wordChainSolverVisitor {
	m := make(map[int][]string, len(startWord))
	for i := 1; i <= len(startWord); i++ {
		m[i] = []string{}
	}
	return &wordChainSolverVisitor{
		wordLength: len(startWord),
		start: startWord,
		end: endWord,
		wordDiff: diffOfWords(startWord, endWord),
		diffCharCountToWords: m,
	}
}

func (v *wordChainSolverVisitor) VisitRow(row []string) {
	if len(row) != 1 {
		return
	}
	if len(row[0]) != v.wordLength {
		return
	}
	diffS := diffOfWords(row[0], v.start)
	diffE := diffOfWords(row[0], v.end)
	if diffS == 0 {
		return
	}
	if diffE == 0 {
		return
	}
	if diffE >= v.wordDiff {
		return
	}
	v.diffCharCountToWords[diffS] = append(v.diffCharCountToWords[diffS], row[0])
}

func diffOfWords(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}
