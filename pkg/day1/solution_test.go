package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sample_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	maxCal, _ := solution(lines)
	assert.Equal(t, 24000, maxCal[0])
}

func Test_solution_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	maxCal, _ := solution(lines)
	fmt.Printf("Final solution for day 1: %d\n", maxCal[0])
}

func Test_sample_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	_, sumOfTop3 := solution(lines)
	assert.Equal(t, 45000, sumOfTop3)
}

func Test_solution_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	_, sumOfTop3 := solution(lines)
	fmt.Printf("Final solution for day 1: %d\n", sumOfTop3)
}
