package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./3/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(ProcessInput(lines))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func ProcessInput(lines []string) int {
	var (
		a     string
		l     string
		b     string
		rowNo int
		tot   int
	)

	for _, line := range lines {
		a = l
		l = b
		b = line
		sum := ProcessLine(a, l, b)
		tot += sum
		rowNo += 1
		fmt.Println(rowNo, sum, tot)
	}

	a = l
	l = b
	b = ""
	sum := ProcessLine(a, l, b)
	tot += sum
	rowNo += 1
	fmt.Println(rowNo, sum, tot)

	return tot
}

func ProcessLine(above string, line string, below string) int {
	if above == "" {
		for i := 0; i < len(line); i++ {
			above = fmt.Sprintf("%s%s", above, ".")
		}
	}
	if below == "" {
		for i := 0; i < len(line); i++ {
			below = fmt.Sprintf("%s%s", above, ".")
		}
	}

	var (
		idx   int
		found bool

		gearIdx int
		sum     int
	)

	found = true
	idx = 0
	for {
		gearIdx, found = getNextGear(line, idx)
		if !found {
			break
		}

		var nums []int
		nums = append(nums, processGearIdx(gearIdx, above)...)
		nums = append(nums, processGearIdx(gearIdx, line)...)
		nums = append(nums, processGearIdx(gearIdx, below)...)

		if len(nums) > 2 {
			fmt.Println("bigger")
		}

		gearSum := 0
		if len(nums) > 1 {
			gearSum = 1
			for _, num := range nums {
				gearSum *= num
			}
		}

		sum += gearSum

		idx = gearIdx + 1
	}

	return sum
}

func processGearIdx(gearIdx int, line string) []int {
	var nums []int
	idx := 0
	for {
		startIdx, endIdx, number, found := GetNextNumber(line, idx)
		if !found {
			break
		}

		if indexMatch(gearIdx, startIdx, endIdx) {
			nums = append(nums, number)
		}
		idx = endIdx + 1
	}

	return nums
}

func indexMatch(idx int, startIdx int, endIdx int) bool {
	if startIdx <= idx-1 && endIdx >= idx-1 {
		return true
	}
	if startIdx <= idx+1 && endIdx >= idx+1 {
		return true
	}

	return false
}

func isSymbol(txt string) bool {
	return txt == "*"
}

func getNextGear(line string, idx int) (int, bool) {
	for it, ch := range line[idx:] {
		i := it + idx
		if isSymbol(string(ch)) {
			return i, true
		}
	}

	return 0, false
}

func GetNextNumber(line string, idx int) (startIdx int, endIdx int, number int, found bool) {
	var (
		numStr string
		err    error
	)

	for it, ch := range line[idx:] {
		i := it + idx
		_, err := strconv.Atoi(string(ch))
		if err != nil {
			if found {
				break
			} else {
				continue
			}
		}

		if !found {
			startIdx = i
			found = true
		}

		endIdx = i

		numStr = fmt.Sprintf("%s%s", numStr, string(ch))
	}

	if found {
		number, err = strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
	}

	return startIdx, endIdx, number, found
}
