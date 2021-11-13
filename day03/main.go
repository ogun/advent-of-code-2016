package main

import (
	"sort"
)

func Part1() int {
	var possibleTriangles int

	for _, sides := range Data {
		t := append([]int{}, sides...)
		sort.Ints(t)

		if t[0]+t[1] > t[2] {
			possibleTriangles++
		}
	}
	return possibleTriangles
}

func Part2() int {
	var possibleTriangles int

	for i := 0; i < len(Data)-2; i += 3 {
		for j := 0; j < 3; j++ {
			vertical := []int{Data[i][j], Data[i+1][j], Data[i+2][j]}
			sort.Ints(vertical)

			if vertical[0]+vertical[1] > vertical[2] {
				possibleTriangles++
			}
		}
	}
	return possibleTriangles
}
