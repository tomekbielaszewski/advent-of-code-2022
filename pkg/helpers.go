package pkg

import (
	"bufio"
	"os"
	"strconv"
)

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func StringToIntNoErr(str string) int {
	atoi, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return atoi
}

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReplaceAtIndex(str string, r rune, i int) string {
	if i+1 > len(str) || i < 0 {
		return str
	}
	out := []rune(str)
	out[i] = r
	return string(out)
}

func SplitAfterNRunes(str string, divEveryRunes int) []string {
	var splitStrings []string
	for len(str) > 0 {
		if len(str) < divEveryRunes {
			divEveryRunes = len(str)
		}
		splitStrings = append(splitStrings, str[:divEveryRunes])
		str = str[divEveryRunes:]
	}
	return splitStrings
}
