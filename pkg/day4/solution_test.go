package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_sample_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 2, sum)
}

func Test_solution_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	fmt.Printf("Final solution for day 4: %d\n", sum)
}

func Test_sample_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	assert.Equal(t, 4, sum)
}

func Test_sample2_second_part(t *testing.T) {
	lines := strings.Split(`4-6,6-8
4-6,5-7
4-6,4-6
4-6,3-5
4-6,2-4
4-6,6-6
4-6,5-5
4-6,4-4
4-6,3-7`, "\n")
	sum := solution2(lines)
	assert.Equal(t, 9, sum)
}

func Test_solution_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	fmt.Printf("Final solution for day 4: %d\n", sum)
}
