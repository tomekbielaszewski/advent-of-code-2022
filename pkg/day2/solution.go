package day1

import "advent-of-code-2022/pkg"

// A X - Rock     - 1
// B Y - Paper    - 2
// C Z - Scissors - 3

func solution(lines []string) int {
	win := 6
	draw := 3
	loose := 0
	rock := 1
	paper := 2
	scissors := 3

	scores := map[string]int{
		"A X": rock + draw,
		"B X": rock + loose,
		"C X": rock + win,
		"A Y": paper + win,
		"B Y": paper + draw,
		"C Y": paper + loose,
		"A Z": scissors + loose,
		"B Z": scissors + win,
		"C Z": scissors + draw,
	}

	sum := 0

	for _, line := range lines {
		sum += scores[line]
	}

	return sum
}

// X - loose
// Y - draw
// Z - win

func solution2(lines []string) int {
	win := 6
	draw := 3
	loose := 0
	rock := 1
	paper := 2
	scissors := 3

	strategy := map[string]rune{
		"A X": 'C',
		"B X": 'A',
		"C X": 'B',
		"A Y": 'A',
		"B Y": 'B',
		"C Y": 'C',
		"A Z": 'B',
		"B Z": 'C',
		"C Z": 'A',
	}

	scores := map[string]int{
		"A A": rock + draw,
		"B A": rock + loose,
		"C A": rock + win,
		"A B": paper + win,
		"B B": paper + draw,
		"C B": paper + loose,
		"A C": scissors + loose,
		"B C": scissors + win,
		"C C": scissors + draw,
	}

	sum := 0

	for _, line := range lines {
		line = pkg.ReplaceAtIndex(line, strategy[line], 2)
		sum += scores[line]
	}

	return sum
}
