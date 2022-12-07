package day1

func solution(lines []string) int {
	runes := []rune(lines[0])

	for i, _ := range runes {
		if i < 4 {
			continue
		}
		uniqueChecker := make(map[rune]int)
		for _, r2 := range runes[i-4 : i] {
			uniqueChecker[r2] = 0
		}
		if len(uniqueChecker) == 4 {
			return i
		}
	}
	return len(runes)
}

func solution2(lines []string) int {
	return 0
}
