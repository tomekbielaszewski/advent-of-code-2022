package day1

func solution(lines []string) int {
	return findEndIndexOfNDistinctRunes(lines[0], 4)
}

func solution2(lines []string) int {
	return findEndIndexOfNDistinctRunes(lines[0], 14)
}

func findEndIndexOfNDistinctRunes(line string, amount int) int {
	runes := []rune(line)

	for i, _ := range runes {
		if i < amount {
			continue
		}
		uniqueChecker := make(map[rune]int)
		for _, r2 := range runes[i-amount : i] {
			uniqueChecker[r2] = 0
		}
		if len(uniqueChecker) == amount {
			return i
		}
	}
	return len(runes)
}
