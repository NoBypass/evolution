package environment

import (
	"evolution/internal/utils"
)

type Environment struct {
	Size      int
	Organisms []*Organism
}

func New(size, maxPop int) *Environment {
	env := &Environment{
		Size:      size,
		Organisms: make([]*Organism, maxPop),
	}

	for i := range env.Organisms {
		env.Organisms[i] = NewOrganism(10, 3, env)
	}

	coords := utils.GenerateUniqueCoordinates(maxPop, size, size)

	for i, org := range env.Organisms {
		org.X = coords[i][0]
		org.Y = coords[i][1]
	}

	return env
}

func (e *Environment) IsOccupied(x, y int) bool {
	for _, org := range e.Organisms {
		if org.X == x && org.Y == y {
			return true
		}
	}
	return false
}
