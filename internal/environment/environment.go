package environment

import (
	"evolution/internal/environment/neural"
	"evolution/internal/utils"
	"math/rand"
)

type SelectionFn func(*Organism) bool

type Environment struct {
	maxPop      int
	Size        int
	selectionFn SelectionFn

	Organisms []*Organism

	CurrGen    int
	CurrGenAge int
	MaxGenAge  int
}

func New(size, maxPop, maxGenAge int, selectionFn SelectionFn) *Environment {
	env := &Environment{
		Size:        size,
		maxPop:      maxPop,
		MaxGenAge:   maxGenAge,
		selectionFn: selectionFn,
		Organisms:   make([]*Organism, maxPop),
	}

	for i := range env.Organisms {
		env.Organisms[i] = NewOrganism(10, 3, env)
	}

	env.RandomizeOrganisms()
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

// GenerateOffspring generates new organisms from the current population
// until the population reaches the maximum population size. The
// mutationRate is used as a quotient of 1 and the mutationRate
func (e *Environment) GenerateOffspring(mutationRate int) {
	for range e.maxPop - len(e.Organisms) {
		offspring := neural.OffspringOf(
			e.Organisms[rand.Intn(len(e.Organisms))].EncodedNet,
			e.Organisms[rand.Intn(len(e.Organisms))].EncodedNet)

		if rand.Intn(mutationRate) == 0 {
			// TODO assure no invalid nets are created: offspring.Mutate()
		}

		e.Organisms = append(e.Organisms, NewOrganismFromNetwork(offspring.Decode(), e))
	}
}

func (e *Environment) RandomizeOrganisms() {
	coords := utils.GenerateUniqueCoordinates(e.maxPop, e.Size, e.Size)

	for i, org := range e.Organisms {
		org.X = coords[i][0]
		org.Y = coords[i][1]
	}
}

func (e *Environment) ApplySelection() {
	for i := len(e.Organisms) - 1; i >= 0; i-- {
		org := e.Organisms[i]
		selected := e.selectionFn(org)
		if selected {
			e.Organisms = append(e.Organisms[:i], e.Organisms[i+1:]...)
		}
	}
}
