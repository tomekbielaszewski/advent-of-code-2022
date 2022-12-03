package day1

import "unicode"

func solution(lines []string) int {
	sum := 0
	for _, line := range lines {
		runes := []rune(line)
		middle := len(line) / 2
		items := make(map[rune]int)
		for i, r := range runes {
			if i < middle {
				diff := 0
				if unicode.IsUpper(r) {
					diff = 38
				} else {
					diff = 96
				}
				items[r] = int(r) - diff
			} else {
				if val, ok := items[r]; ok {
					sum += val
					items[r] = 0
				}
			}
		}
	}
	return sum
}

func solution2(lines []string) int {
	return 0
}
