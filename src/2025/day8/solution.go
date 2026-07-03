package main

import (
	"cmp"
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

	mapOfPoints := make([]Point, len(lines)) // []Point{}

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
		mapOfPoints[index] = Point{index, 0, x, y, z}
	}

	distances := make([][]float64, len(mapOfPoints))

	for i := 0; i < len(mapOfPoints); i++ {
		distances[i] = make([]float64, len(mapOfPoints))
		for j := i + 1; j < len(mapOfPoints); j++ {
			distances[i][j] = math.Sqrt(math.Pow(float64(mapOfPoints[i].x-mapOfPoints[j].x), 2) + math.Pow(float64(mapOfPoints[i].y-mapOfPoints[j].y), 2) + math.Pow(float64(mapOfPoints[i].z-mapOfPoints[j].z), 2))
		}
	}

	for i := 0; i < times; i++ {
		closestPoints := getClosestPoints(&mapOfPoints, &distances)

		if mapOfPoints[closestPoints[0]].chain != 0 && mapOfPoints[closestPoints[1]].chain == 0 {
			mapOfPoints[closestPoints[1]].chain = mapOfPoints[closestPoints[0]].chain
		} else if mapOfPoints[closestPoints[1]].chain != 0 && mapOfPoints[closestPoints[0]].chain == 0 {
			mapOfPoints[closestPoints[0]].chain = mapOfPoints[closestPoints[1]].chain
		} else if mapOfPoints[closestPoints[1]].chain == 0 && mapOfPoints[closestPoints[0]].chain == 0 {
			mapOfPoints[closestPoints[0]].chain = i + 1
			mapOfPoints[closestPoints[1]].chain = i + 1
		}

		// fmt.Println("Closest points:", mapOfPoints[closestPoints[0]].id, mapOfPoints[closestPoints[0]].chain, mapOfPoints[closestPoints[1]].id, mapOfPoints[closestPoints[1]].chain)
	}

	chains := map[int]int{}
	for i := 0; i < len(mapOfPoints); i++ {
		val, exists := chains[mapOfPoints[i].chain]
		if exists {
			chains[mapOfPoints[i].chain] = val + 1
		} else {
			chains[mapOfPoints[i].chain] = 1
		}
	}

	keys := make([]int, 0, len(chains))
	for k, v := range chains {
		if k == 0 {
			continue // chains[0] - это остатки, не использовать в перемножении
		}
		keys = append(keys, v)
	}
	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	result := 1
	for i := 0; i < 3; i++ {
		fmt.Println(keys[i])
		result *= keys[i]
	}

	fmt.Println("Ответ:", result)
}

func getClosestPoints(mapOfPoints *[]Point, distances *[][]float64) [2]int {
	distance := math.MaxFloat64
	curI := 0
	curJ := 0
	for i := 0; i < len(*distances); i++ {
		for j := i + 1; j < len((*distances)[i]); j++ {
			// fmt.Println(distance, (*distances)[i][j])
			if distance > (*distances)[i][j] && ((*mapOfPoints)[i].chain == 0 || (*mapOfPoints)[j].chain == 0) {
				distance = (*distances)[i][j]
				curI = i
				curJ = j
			}
		}
	}

	fmt.Println("Min:", (*mapOfPoints)[curI].x, (*mapOfPoints)[curI].y, (*mapOfPoints)[curI].z, "-", (*mapOfPoints)[curJ].x, (*mapOfPoints)[curJ].y, (*mapOfPoints)[curJ].z)

	return [2]int{curI, curJ}
}
