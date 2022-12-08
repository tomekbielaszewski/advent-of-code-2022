package day1

func solution(lines []string) int {
	var forest [][]rune
	for _, line := range lines {
		forest = append(forest, []rune(line))
	}

	sum := 0

	for y, treeLine := range forest {
		for x, tree := range treeLine {
			if isVisible(x, y, tree, forest) {
				sum++
			}
		}
	}

	return sum
}

func isVisible(x int, y int, tree rune, forest [][]rune) bool {
	visible := true

	for i := 0; i < y; i++ {
		visible = visible && (forest[i][x] < tree)
	}
	if visible {
		return visible
	}

	visible = true
	for i := len(forest) - 1; i > y; i-- {
		visible = visible && (forest[i][x] < tree)
	}
	if visible {
		return visible
	}

	visible = true
	for i := 0; i < x; i++ {
		visible = visible && (forest[y][i] < tree)
	}
	if visible {
		return visible
	}

	visible = true
	for i := len(forest[0]) - 1; i > x; i-- {
		visible = visible && (forest[y][i] < tree)
	}
	return visible
}

func solution2(lines []string) int {
	return 0
}
