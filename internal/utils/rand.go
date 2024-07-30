package utils

import (
	"math/rand"
)

func GenerateUniqueCoordinates(count, maxX, maxY int) [][2]int {
	coordinates := make(map[[2]int]struct{})
	result := make([][2]int, 0, count)

	for len(result) < count {
		x := rand.Intn(maxX)
		y := rand.Intn(maxY)
		coord := [2]int{x, y}
		if _, exists := coordinates[coord]; !exists {
			coordinates[coord] = struct{}{}
			result = append(result, coord)
		}
	}

	return result
}
