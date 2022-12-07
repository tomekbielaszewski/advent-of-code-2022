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
	assert.Equal(t, 7, sum)
}

func Test_sample2_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample2.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 5, sum)
}

func Test_sample3_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample3.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 6, sum)
}

func Test_sample4_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample4.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 10, sum)
}

func Test_sample5_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample5.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 11, sum)
}

func Test_solution_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	fmt.Printf("Final solution for day 6: %d\n", sum)
}

func Test_sample_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	assert.Equal(t, "MCD", sum)
}

func Test_solution_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	fmt.Printf("Final solution for day 6: %d\n", sum)
}
