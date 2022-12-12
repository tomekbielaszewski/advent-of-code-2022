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
	sum := solution(lines)
	assert.Equal(t, 10605, sum)
}

func Test_solution_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	fmt.Printf("Final solution for day 11: %d\n", sum)
}

func Test_sample_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	assert.Equal(t, 2713310158, sum)
}

func Test_solution_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	fmt.Printf("Final solution for day 11: %d\n", sum)
}
