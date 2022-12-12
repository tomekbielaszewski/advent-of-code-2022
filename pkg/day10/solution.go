package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"math"
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

func solution2(lines []string) string {
	screen := make([][]rune, 6)
	register := 1
	cycle := 0

	for i := 0; i < len(screen); i++ {
		screen[i] = []rune("........................................")
	}
	draw(screen, register, cycle)

	for _, line := range lines {
		split := strings.Split(line, " ")
		switch split[0] {
		case "noop":
			cycle++
			draw(screen, register, cycle)
			break
		case "addx":
			cycle++
			draw(screen, register, cycle)
			cycle++
			register += pkg.StringToIntNoErr(split[1])
			draw(screen, register, cycle)
			break
		}
	}

	result := ""

	for _, runes := range screen {
		result = fmt.Sprintf("%s\n%s", result, string(runes))
	}

	return result
}

func draw(screen [][]rune, register int, cycle int) {
	y := cycle / 40
	x := cycle % 40
	if math.Abs(float64(x-register)) <= 1 {
		screen[y][x] = '#'
	}
}

func signalStrength(reg int, cycle int, interestingValues map[int]int) {
	if (cycle+20)%40 == 0 {
		interestingValues[cycle] = reg * cycle
	}
}
