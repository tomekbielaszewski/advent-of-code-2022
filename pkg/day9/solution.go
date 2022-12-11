package day1

import (
	"advent-of-code-2022/pkg"
	"fmt"
	"math"
	"strings"
)

func solution(lines []string) int {
	head := &point{}
	tail := &point{}
	tailLocations := make(map[string]int)
	tailLocations[tail.ToString()] = 0

	for _, line := range lines {
		split := strings.Split(line, " ")
		dir := split[0]
		steps := pkg.StringToIntNoErr(split[1])

		for i := 0; i < steps; i++ {
			head.move(dir)
			tail.follow(head)
			tailLocations[tail.ToString()] = 0
		}
	}

	return len(tailLocations)
}

func solution2(lines []string) int {
	return 0
}

type point struct {
	x, y int
	old  *point
}

func (p1 *point) move(dir string) {
	p1.old = &point{
		x:   p1.x,
		y:   p1.y,
		old: nil,
	}
	switch dir {
	case "R":
		p1.x++
		break
	case "L":
		p1.x--
		break
	case "U":
		p1.y--
		break
	case "D":
		p1.y++
		break
	}
}

func (p1 *point) follow(p2 *point) {
	if p1.sqDistance(p2) >= 4 {
		p1.set(p2.old)
	}
}

func (p1 *point) sqDistance(p2 *point) float64 {
	return math.Pow(float64(p1.x-p2.x), 2) +
		math.Pow(float64(p1.y-p2.y), 2)
}

func (p1 *point) ToString() string {
	return fmt.Sprintf("%d%d", p1.x, p1.y)
}

func (p1 *point) set(to *point) {
	if to == nil {
		return
	}
	p1.x = to.x
	p1.y = to.y
}
