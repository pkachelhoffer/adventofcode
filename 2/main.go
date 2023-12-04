package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube string

var (
	CubeRed   Cube = "red"
	CubeBlue  Cube = "blue"
	CubeGreen Cube = "green"
)

type Game struct {
	Number int
	Sets   []Set
}

type Set struct {
	Cubes []CubePlay
}

type CubePlay struct {
	Cube  Cube
	Count int
}

func (g Game) IsPossible(limits []CubePlay) bool {
	for _, set := range g.Sets {
		countMap := make(map[Cube]int)

		for _, cube := range set.Cubes {
			countMap[cube.Cube] += cube.Count
		}

		for key, val := range countMap {
			for _, limit := range limits {
				if limit.Cube == key && val > limit.Count {
					return false
				}
			}
		}
	}

	return true
}

func (g Game) Limit() []CubePlay {
	countMap := make(map[Cube]int)

	for _, set := range g.Sets {
		for _, cube := range set.Cubes {
			if cube.Count > countMap[cube.Cube] {
				countMap[cube.Cube] = cube.Count
			}
		}
	}

	var limit []CubePlay
	for key, val := range countMap {
		limit = append(limit, CubePlay{
			Cube:  key,
			Count: val,
		})
	}

	return limit
}

func main() {
	file, err := os.Open("./2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		game := parseGame(scanner.Text())

		fewest := game.Limit()

		pow := fewest[0].Count
		for i := 1; i < len(fewest); i++ {
			pow = pow * fewest[i].Count
		}

		sum += pow
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func parseGame(line string) Game {
	split := strings.Split(line, ":")
	gameTitleStr := split[0]
	gamePlaysStr := split[1]

	gameTitleSplit := strings.Split(gameTitleStr, " ")

	gameNumber, err := strconv.Atoi(gameTitleSplit[1])
	if err != nil {
		panic(err)
	}

	sets := parseGamePlay(gamePlaysStr)

	return Game{
		Number: gameNumber,
		Sets:   sets,
	}
}

// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGamePlay(gamePlay string) []Set {
	var sets []Set
	gameSplit := strings.Split(gamePlay, ";")
	for _, setStr := range gameSplit {
		set := Set{
			Cubes: parseSet(setStr),
		}
		sets = append(sets, set)
	}

	return sets
}

// 3 blue, 4 red
func parseSet(set string) []CubePlay {
	var cubes []CubePlay

	setSplit := strings.Split(set, ",")
	for _, cubeStr := range setSplit {
		cube, count := parseCube(cubeStr)
		cubes = append(cubes, CubePlay{
			Cube:  cube,
			Count: count,
		})
	}

	return cubes
}

// 3 blue
func parseCube(cube string) (Cube, int) {
	split := strings.Split(strings.TrimSpace(cube), " ")
	count, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	return Cube(split[1]), count
}
