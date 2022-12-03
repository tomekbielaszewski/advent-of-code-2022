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
				items[r] = itemPriority(r)
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
	sum := 0
	items := make(map[rune]int)
	for i, line := range lines {
		runes := []rune(line)
		if i%3 == 0 {
			items = make(map[rune]int)
			for _, r := range runes {
				items[r] = itemPriority(r)
			}
		} else {
			repeatedItems := make(map[rune]int)
			for _, r := range runes {
				if val, ok := items[r]; ok {
					repeatedItems[r] = val
				}
			}
			items = repeatedItems
			if i%3 == 2 {
				for _, priority := range items {
					sum += priority
				}
			}
		}
	}
	return sum
}

func itemPriority(r rune) int {
	diff := 0
	if unicode.IsUpper(r) {
		diff = 38
	} else {
		diff = 96
	}
	return int(r) - diff
}
