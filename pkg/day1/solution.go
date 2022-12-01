package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"sort"
)

func solution(lines []string) ([]int, int) {
	var sums []int
	currentSum := 0
	for _, line := range lines {
		cal, err := pkg.StringToInt(line)
		if err != nil {
			sums = append(sums, currentSum)
			currentSum = 0
		} else {
			currentSum += cal
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	fmt.Printf("Max calories is carried by elf with %d calories at hands\n", sums[0])
	return sums, sums[0] + sums[1] + sums[2]
}
