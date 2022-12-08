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

func solution2(lines []string) int {
	var forest [][]rune
	for _, line := range lines {
		forest = append(forest, []rune(line))
	}

	score := 1

	for y, treeLine := range forest {
		for x, tree := range treeLine {
			thisScore := calcScenicScore(x, y, tree, forest)
			if thisScore > score {
				score = thisScore
			}
		}
	}

	return score
}

func calcScenicScore(x int, y int, tree rune, forest [][]rune) int {
	northScore := 0
	for i := y; i >= 0; i-- {
		if y != i && forest[i][x] < tree {
			northScore++
		}
		if y != i && forest[i][x] >= tree {
			northScore++
			break
		}
	}

	southScore := 0
	for i := y; i < len(forest); i++ {
		if y != i && forest[i][x] < tree {
			southScore++
		}
		if y != i && forest[i][x] >= tree {
			southScore++
			break
		}
	}

	westScore := 0
	for i := x; i >= 0; i-- {
		if x != i && forest[y][i] < tree {
			westScore++
		}
		if x != i && forest[y][i] >= tree {
			westScore++
			break
		}
	}

	eastScore := 0
	for i := x; i < len(forest[0]); i++ {
		if x != i && forest[y][i] < tree {
			eastScore++
		}
		if x != i && forest[y][i] >= tree {
			eastScore++
			break
		}
	}

	return northScore * southScore * eastScore * westScore
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
