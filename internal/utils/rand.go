package utils

import (
	"log"
	"math/rand"
)

// GenerateUniqueCoordinates generates a slice of unique coordinates
// within the range of maxX and maxY. The count parameter specifies
// the number of unique coordinates to generate.
func GenerateUniqueCoordinates(count, maxX, maxY int) [][2]int {
	if count > maxX*maxY {
		log.Fatal("count is greater than maxX*maxY")
	}

	allCoords := make([][2]int, 0, maxX*maxY)
	for x := range maxX {
		for y := range maxY {
			allCoords = append(allCoords, [2]int{x, y})
		}
	}

	rand.Shuffle(len(allCoords), func(i, j int) {
		allCoords[i], allCoords[j] = allCoords[j], allCoords[i]
	})

	return allCoords[:count]
}
