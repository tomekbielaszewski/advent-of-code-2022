package day1

import (
	"advent-of-code-2022/pkg"
	"strconv"
	"strings"
)

func solution(lines []string) string {
	stacksInitialized := false
	var stackLines []string
	var moveLines []string
	for _, line := range lines {
		if line == "" {
			stacksInitialized = true
			continue
		}

		if !stacksInitialized {
			stackLines = append(stackLines, line)
		} else {
			moveLines = append(moveLines, line)
		}
	}

	stacks := parseStacks(stackLines)

	executeMovingCrates(moveLines, stacks)

	return peekStackTops(stacks)
}

func solution2(lines []string) string {
	stacksInitialized := false
	var stackLines []string
	var moveLines []string
	for _, line := range lines {
		if line == "" {
			stacksInitialized = true
			continue
		}

		if !stacksInitialized {
			stackLines = append(stackLines, line)
		} else {
			moveLines = append(moveLines, line)
		}
	}

	stacks := parseStacks(stackLines)

	executeMovingCrates2(moveLines, stacks)

	return peekStackTops(stacks)
}

type stack struct {
	crates []string
}

func parseStacks(lines []string) map[string]*stack {
	stacks := make(map[string]*stack)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if i == len(lines)-1 {
			split := strings.Split(line, "   ")
			for _, s := range split {
				id := strings.TrimSpace(s)
				stacks[id] = &stack{}
			}
		} else {
			crates := pkg.SplitAfterNRunes(line, 4)
			for i, crate := range crates {
				crate = strings.TrimSpace(crate)
				if crate != "" {
					stacks[strconv.Itoa(i+1)].crates = append(stacks[strconv.Itoa(i+1)].crates, crate)
				}
			}
		}
	}
	return stacks
}

func executeMovingCrates(lines []string, stacks map[string]*stack) {
	for _, line := range lines {
		params := strings.Split(line, " ")
		stacks[params[3]].moveCrates(pkg.StringToIntNoErr(params[1]), stacks[params[5]])
	}
}

func executeMovingCrates2(lines []string, stacks map[string]*stack) {
	for _, line := range lines {
		params := strings.Split(line, " ")
		stacks[params[3]].moveCrates2(pkg.StringToIntNoErr(params[1]), stacks[params[5]])
	}
}

func peekStackTops(stacks map[string]*stack) string {
	stackTops := ""
	for i := 1; i < len(stacks)+1; i++ {
		crates := stacks[strconv.Itoa(i)].crates
		stackTops += crates[len(crates)-1]
	}

	stackTops = strings.Replace(stackTops, "[", "", -1)
	stackTops = strings.Replace(stackTops, "]", "", -1)

	return stackTops
}

func (from *stack) moveCrates(num int, to *stack) {
	for i := 0; i < num; i++ {
		to.crates = append(to.crates, from.crates[len(from.crates)-1])
		from.crates = from.crates[:len(from.crates)-1]
	}
}

func (from *stack) moveCrates2(num int, to *stack) {
	to.crates = append(to.crates, from.crates[len(from.crates)-num:]...)
	from.crates = from.crates[:len(from.crates)-num]
}
