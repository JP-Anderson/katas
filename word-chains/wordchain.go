package wordchain

func Solve(start, end string) []string {
	if len(start) != len(end) {
		return []string{}
	}
	return nil
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
