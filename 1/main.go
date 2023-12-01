package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSum := processLine(line)

		total += lineSum

		fmt.Println(line, lineSum, total)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(total)
}

var words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func processLine(line string) int {
	first := 0
	last := 0
	isFirst := true
	for i, ch := range line {
		var (
			num int
			ok  bool
		)

		convNum, err := strconv.Atoi(string(ch))
		if err != nil {
			num, ok = convertToNumber(line[i:])
			if !ok {
				continue
			}
		} else {
			num = convNum
		}

		if isFirst {
			first = num
			isFirst = false
		}
		last = num
	}

	ret, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	if err != nil {
		panic(err)
	}

	return ret
}

func convertToNumber(line string) (int, bool) {
	for i, word := range words {
		if len(line) < len(word) {
			continue
		}

		reject := false
		for i := 0; i < len(word); i++ {
			if word[i] != line[i] {
				reject = true
				break
			}
		}

		if !reject {
			return i, true
		}
	}

	return 0, false
}
