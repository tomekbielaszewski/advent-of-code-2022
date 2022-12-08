package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runes_and_ascii_code_for_numbers(t *testing.T) {
	numbers := []rune{
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
	}
	for i, number := range numbers {
		if i == 0 {
			continue
		}
		assert.True(t, numbers[i-1] < number)
	}
}

func Test_sample_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	assert.Equal(t, 21, sum)
}

func Test_solution_first_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution(lines)
	fmt.Printf("Final solution for day 8: %d\n", sum)
}

func Test_sample_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("sample.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	assert.Equal(t, 8, sum)
}

func Test_solution_second_part(t *testing.T) {
	lines, err := pkg.ReadAllLines("final.txt")
	if err != nil {
		t.FailNow()
		return
	}
	sum := solution2(lines)
	fmt.Printf("Final solution for day 8: %d\n", sum)
}
