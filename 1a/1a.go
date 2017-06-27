package main

import (
	"fmt"
	"strconv"
	"strings"
)

// pos {direction, x, y}
// Direction: 0 is North, 1 is East, 2 is South, 3 is West
// Coordinates: Steps North/East of starting point
func main() {
	input := strings.Split("L3, R1, L4, L1, L2, R4, L3, L3, R2, R3, L5, R1, R3, L4, L1, L2, R2, R1, L4, L4, R2, L5, R3, R2, R1, L1, L2, R2, R2, L1, L1, R2, R1, L3, L5, R4, L3, R3, R3, L5, L190, L4, R4, R51, L4, R5, R5, R2, L1, L3, R1, R4, L3, R1, R3, L5, L4, R2, R5, R2, L1, L5, L1, L1, R78, L3, R2, L3, R5, L2, R2, R4, L1, L4, R1, R185, R3, L4, L1, L1, L3, R4, L4, L1, R5, L5, L1, R5, L1, R2, L5, L2, R4, R3, L2, R3, R1, L3, L5, L4, R3, L2, L4, L5, L4, R1, L1, R5, L2, R4, R2, R3, L1, L1, L4, L3, R4, L3, L5, R2, L5, L1, L1, R2, R3, L5, L3, L2, L1, L4, R4, R4, L2, R3, R1, L2, R1, L2, L2, R3, R3, L1, R4, L5, L3, R4, R4, R1, L2, L5, L3, R1, R4, L2, R5, R4, R2, L5, L3, R4, R1, L1, R5, L3, R1, R5, L2, R1, L5, L2, R2, L2, L3, R3, R3, R1", ", ")
	pos := [3]int{0, 0, 0}

	for _, move := range input {
		pos = turn(pos, move)
		steps, _ := strconv.Atoi(move[1:len(move)])
		for i := 0; i < steps; i++ {
			pos = step(pos)
		}
	}

	// Get absolute values and print the distance from the starting point
	distance := 0
	for i := 1; i < 3; i++ {
		if pos[i] < 0 {
			distance -= pos[i]
		} else {
			distance += pos[i]
		}
	}
	fmt.Println(distance)
}

func turn(pos [3]int, move string) [3]int {
	switch move[0] {
	case 'L':
		pos[0] = (pos[0] + 3) % 4
	case 'R':
		pos[0] = (pos[0] + 1) % 4
	}
	return pos
}

func step(pos [3]int) [3]int {
	switch pos[0] {
	case 0:
		pos[2]++
	case 1:
		pos[1]++
	case 2:
		pos[2]--
	case 3:
		pos[1]--
	}
	return pos
}
