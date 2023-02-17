package wordchain

func Solve(start, end string) []string {
	if len(start) != len(end) {
		return []string{}
	}
	
	return nil
}

type wordChainSolverVisitor struct {
	wordLength int
	start string
	diffCharCountToWords map[int][]string
}

func newVisitor(startWord string) *wordChainSolverVisitor {
	m := make(map[int][]string, len(startWord))
	for i := 1; i <= len(startWord); i++ {
		m[i] = []string{}
	}
	return &wordChainSolverVisitor{
		wordLength: len(startWord),
		start: startWord,
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
	diff := diffOfWords(row[0], v.start)
	if diff == 0 {
		return
	}
	v.diffCharCountToWords[diff] = append(v.diffCharCountToWords[diff], row[0])
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
