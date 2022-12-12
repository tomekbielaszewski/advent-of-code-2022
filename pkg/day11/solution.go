package day1

import (
	"advent-of-code-2022/pkg"
	"sort"
	"strings"
)

func solution(lines []string) int {
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.inspectItems(monkeys)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount > monkeys[j].inspectionCount
	})

	return monkeys[0].inspectionCount * monkeys[1].inspectionCount
}

func solution2(lines []string) int {
	monkeys := parseMonkeys(lines)
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.inspectItems2(monkeys)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount > monkeys[j].inspectionCount
	})

	return monkeys[0].inspectionCount * monkeys[1].inspectionCount
}

func parseMonkeys(lines []string) []*monkey {
	var monkeys []*monkey
	for i := 0; i < len(lines); i += 7 {
		itemsStr := strings.ReplaceAll(lines[i+1], "  Starting items: ", "")
		itemsStrArr := strings.Split(itemsStr, ", ")
		items := make([]int, len(itemsStrArr))
		for i2, s := range itemsStrArr {
			items[i2] = pkg.StringToIntNoErr(s)
		}

		operationMul := "1"
		operationAdd := "0"
		if strings.HasPrefix(lines[i+2], "  Operation: new = old * ") {
			operationMul = strings.ReplaceAll(lines[i+2], "  Operation: new = old * ", "")
		} else {
			operationAdd = strings.ReplaceAll(lines[i+2], "  Operation: new = old + ", "")
		}

		testDivisibleStr := strings.ReplaceAll(lines[i+3], "  Test: divisible by ", "")
		testTrueStr := strings.ReplaceAll(lines[i+4], "    If true: throw to monkey ", "")
		testFalseStr := strings.ReplaceAll(lines[i+5], "    If false: throw to monkey ", "")

		monkeys = append(monkeys, &monkey{
			items:           items,
			operationMul:    operationMul,
			operationAdd:    operationAdd,
			testDivisible:   pkg.StringToIntNoErr(testDivisibleStr),
			testTrue:        pkg.StringToIntNoErr(testTrueStr),
			testFalse:       pkg.StringToIntNoErr(testFalseStr),
			inspectionCount: 0,
		})
	}
	massDiv := 1
	for _, m := range monkeys {
		massDiv = massDiv * m.testDivisible
	}
	for _, m := range monkeys {
		m.massDiv = massDiv
	}
	return monkeys
}

type monkey struct {
	items           []int
	operationMul    string
	operationAdd    string
	testDivisible   int
	testTrue        int
	testFalse       int
	inspectionCount int
	massDiv         int
}

func (m *monkey) inspectItems(monkeys []*monkey) {
	for i := 0; i < len(m.items); i++ {
		var operationMul int
		var operationAdd int

		if m.operationMul == "old" {
			operationMul = m.items[i]
		} else {
			operationMul = pkg.StringToIntNoErr(m.operationMul)
		}

		if m.operationAdd == "old" {
			operationAdd = m.items[i]
		} else {
			operationAdd = pkg.StringToIntNoErr(m.operationAdd)
		}

		m.items[i] = m.items[i] * operationMul
		m.items[i] = m.items[i] + operationAdd
		m.items[i] = m.items[i] / 3
		m.inspectionCount++
		if m.items[i]%m.testDivisible == 0 {
			monkeys[m.testTrue].items = append(monkeys[m.testTrue].items, m.items[i])
		} else {
			monkeys[m.testFalse].items = append(monkeys[m.testFalse].items, m.items[i])
		}
	}
	m.items = nil
}

func (m *monkey) inspectItems2(monkeys []*monkey) {
	for i := 0; i < len(m.items); i++ {
		var operationMul int
		var operationAdd int

		if m.operationMul == "old" {
			operationMul = m.items[i]
		} else {
			operationMul = pkg.StringToIntNoErr(m.operationMul)
		}

		if m.operationAdd == "old" {
			operationAdd = m.items[i]
		} else {
			operationAdd = pkg.StringToIntNoErr(m.operationAdd)
		}

		m.items[i] = m.items[i] * operationMul
		m.items[i] = m.items[i] + operationAdd
		m.inspectionCount++

		m.items[i] = m.items[i] % m.massDiv

		if m.items[i]%m.testDivisible == 0 {
			monkeys[m.testTrue].items = append(monkeys[m.testTrue].items, m.items[i])
		} else {
			monkeys[m.testFalse].items = append(monkeys[m.testFalse].items, m.items[i])
		}
	}
	m.items = nil
}
