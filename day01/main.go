package main

import "strconv"

var Data []string = []string{"R4", "R1", "L2", "R1", "L1", "L1", "R1", "L5", "R1", "R5", "L2", "R3", "L3", "L4", "R4", "R4", "R3", "L5", "L1", "R5", "R3", "L4", "R1", "R5", "L1", "R3", "L2", "R3", "R1", "L4", "L1", "R1", "L1", "L5", "R1", "L2", "R2", "L3", "L5", "R1", "R5", "L1", "R188", "L3", "R2", "R52", "R5", "L3", "R79", "L1", "R5", "R186", "R2", "R1", "L3", "L5", "L2", "R2", "R4", "R5", "R5", "L5", "L4", "R5", "R3", "L4", "R4", "L4", "L4", "R5", "L4", "L3", "L1", "L4", "R1", "R2", "L5", "R3", "L4", "R3", "L3", "L5", "R1", "R1", "L3", "R2", "R1", "R2", "R2", "L4", "R5", "R1", "R3", "R2", "L2", "L2", "L1", "R2", "L1", "L3", "R5", "R1", "R4", "R5", "R2", "R2", "R4", "R4", "R1", "L3", "R4", "L2", "R2", "R1", "R3", "L5", "R5", "R2", "R5", "L1", "R2", "R4", "L1", "R5", "L3", "L3", "R1", "L4", "R2", "L2", "R1", "L1", "R4", "R3", "L2", "L3", "R3", "L2", "R1", "L4", "R5", "L1", "R5", "L2", "L1", "L5", "L2", "L5", "L2", "L4", "L2", "R3"}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func Part1() int {
	rMap := map[[2]int][2]int{
		{0, 0}:  {1, 0},
		{0, 1}:  {-1, 0},
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{-1, 0}: {0, -1},
	}

	lMap := map[[2]int][2]int{
		{0, 0}:  {-1, 0},
		{0, 1}:  {1, 0},
		{0, -1}: {-1, 0},
		{1, 0}:  {0, -1},
		{-1, 0}: {0, 1},
	}

	currentPoint := [2]int{0, 0}
	currentDirection := [2]int{0, 0}
	for _, target := range Data {
		direction := target[0]
		step, _ := strconv.Atoi(target[1:])

		if direction == 'R' {
			currentDirection = rMap[currentDirection]
		} else {
			currentDirection = lMap[currentDirection]
		}

		currentPoint[0] += currentDirection[0] * step
		currentPoint[1] += currentDirection[1] * step
	}

	return abs(currentPoint[0]) + abs(currentPoint[1])
}

func Part2() int {
	rMap := map[[2]int][2]int{
		{0, 0}:  {1, 0},
		{0, 1}:  {-1, 0},
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{-1, 0}: {0, -1},
	}

	lMap := map[[2]int][2]int{
		{0, 0}:  {-1, 0},
		{0, 1}:  {1, 0},
		{0, -1}: {-1, 0},
		{1, 0}:  {0, -1},
		{-1, 0}: {0, 1},
	}

	visited := make(map[[2]int]bool)
	currentPoint := [2]int{0, 0}
	currentDirection := [2]int{0, 0}
	for _, target := range Data {
		direction := target[0]
		step, _ := strconv.Atoi(target[1:])

		if direction == 'R' {
			currentDirection = rMap[currentDirection]
		} else {
			currentDirection = lMap[currentDirection]
		}

		for i := 0; i < step; i++ {
			currentPoint[0] += currentDirection[0]
			currentPoint[1] += currentDirection[1]

			if _, found := visited[currentPoint]; !found {
				visited[currentPoint] = true
			} else {
				return abs(currentPoint[0]) + abs(currentPoint[1])
			}
		}
	}

	return 0
}
