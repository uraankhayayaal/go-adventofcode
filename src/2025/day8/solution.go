package main

import (
	"fmt"
	"go-adventofcode/src/helpers"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	id    int
	chain int // 0 - no chained
	x     int
	y     int
	z     int
}

func Solution(inputFilePath string, mode string) {
	fmt.Println("Start of Solution of task 2025 day 8")

	times := 10
	if mode == "prod" {
		times = 1000
	}

	lines, err := helpers.ReadFileLines(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to read test file: %v", err)
	}

	mapOfPints := make([]Point, len(lines)) // []Point{}

	for index, line := range lines {
		coordinates := strings.Split(line, ",")
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			fmt.Println("Error during conversion:", err)
			return
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			fmt.Println("Error during conversion:", err)
			return
		}
		z, err := strconv.Atoi(coordinates[2])
		if err != nil {
			fmt.Println("Error during conversion:", err)
			return
		}
		mapOfPints[index] = Point{index, 0, x, y, z}
	}

	distances := make([][]float64, len(mapOfPints))

	for i := 0; i < len(mapOfPints); i++ {
		distances[i] = make([]float64, len(mapOfPints))
		for j := i + 1; j < len(mapOfPints); j++ {
			distances[i][j] = math.Sqrt(math.Pow(float64(mapOfPints[i].x-mapOfPints[j].x), 2) + math.Pow(float64(mapOfPints[i].y-mapOfPints[j].y), 2) + math.Pow(float64(mapOfPints[i].z-mapOfPints[j].z), 2))
		}
	}

	for i := 0; i < times; i++ {
		closestPoints := getClosestPoints(&mapOfPints, &distances)

		if mapOfPints[closestPoints[0]].chain != 0 && mapOfPints[closestPoints[1]].chain == 0 {
			mapOfPints[closestPoints[1]].chain = mapOfPints[closestPoints[0]].chain
		}
		if mapOfPints[closestPoints[1]].chain != 0 && mapOfPints[closestPoints[0]].chain == 0 {
			mapOfPints[closestPoints[0]].chain = mapOfPints[closestPoints[1]].chain
		}
		if mapOfPints[closestPoints[1]].chain == 0 && mapOfPints[closestPoints[0]].chain == 0 {
			mapOfPints[closestPoints[0]].chain = i + 1
			mapOfPints[closestPoints[1]].chain = i + 1
		}

		fmt.Println("Closest points:", mapOfPints[closestPoints[0]].id, mapOfPints[closestPoints[0]].chain, mapOfPints[closestPoints[1]].id, mapOfPints[closestPoints[1]].chain)
	}

	chains := map[int]int{}
	for i := 0; i < len(mapOfPints); i++ {
		val, exists := chains[mapOfPints[i].chain]
		if exists {
			chains[mapOfPints[i].chain] = val + 1
		} else {
			chains[mapOfPints[i].chain] = 1
		}
	}

	for i := 0; i < len(chains); i++ {
		fmt.Println("chain:", i, chains[i])
	}

	keys := make([]int, 0, len(chains))
	for k := range chains {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for i := 0; i < len(keys); i++ {
		fmt.Println("keys:", keys[i])
	}
}

func getClosestPoints(mapOfPints *[]Point, distances *[][]float64) [2]int {
	distance := math.MaxFloat64
	curI := 0
	curJ := 0
	for i := 0; i < len(*distances); i++ {
		for j := i + 1; j < len((*distances)[i]); j++ {
			// fmt.Println(distance, (*distances)[i][j])
			if distance > (*distances)[i][j] && ((*mapOfPints)[i].chain == 0 || (*mapOfPints)[j].chain == 0) {
				distance = (*distances)[i][j]
				curI = i
				curJ = j
			}
		}
	}

	return [2]int{curI, curJ}
}
