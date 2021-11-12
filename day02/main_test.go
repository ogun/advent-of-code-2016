package main

import "testing"

func TestPart1(t *testing.T) {
	want := "99332"
	value := Part1()

	if value != want {
		t.Errorf("	() = %v, want %v", value, want)
	}
}

func TestPart2(t *testing.T) {
	want := "DD483"
	value := Part2()

	if value != want {
		t.Errorf("Part2() = %v, want %v", value, want)
	}
}
