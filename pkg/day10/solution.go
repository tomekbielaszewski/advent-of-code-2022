package day1

import (
	"advent-of-code-2022/pkg"
	"strings"
)

func solution(lines []string) int {
	cycle := 1
	reg := 1
	strengths := make(map[int]int)

	for _, line := range lines {
		split := strings.Split(line, " ")
		switch split[0] {
		case "noop":
			cycle++
			signalStrength(reg, cycle, strengths)
			break
		case "addx":
			cycle++
			signalStrength(reg, cycle, strengths)
			cycle++
			reg += pkg.StringToIntNoErr(split[1])
			signalStrength(reg, cycle, strengths)
			break
		}
	}

	sum := 0
	for _, v := range strengths {
		sum += v
	}

	return sum
}

func solution2(lines []string) int {
	return 0
}

func signalStrength(reg int, cycle int, interestingValues map[int]int) {
	if (cycle+20)%40 == 0 {
		interestingValues[cycle] = reg * cycle
	}
}
