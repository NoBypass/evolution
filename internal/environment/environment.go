package environment

import (
	"evolution/internal/environment/neural"
	"evolution/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"time"
)

type SelectionFn func(*Organism) bool

type Environment struct {
	MaxPop      int
	Size        int
	selectionFn SelectionFn

	Organisms []*Organism

	CurrGen    int
	CurrGenAge int
	MaxGenAge  int
	Survivors  int

	Paused   bool
	MSPT     int64
	lastTick time.Time
}

func New(size, maxPop, maxGenAge int, selectionFn SelectionFn) *Environment {
	env := &Environment{
		Size:        size,
		MaxPop:      maxPop,
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
// until the population reaches the maximum population size.
func (e *Environment) GenerateOffspring(mutationRate int) {
	for range e.MaxPop - len(e.Organisms) {
		offspring := neural.OffspringOf(
			e.Organisms[rand.Intn(len(e.Organisms))].EncodedNet,
			e.Organisms[rand.Intn(len(e.Organisms))].EncodedNet)

		if rand.Intn(mutationRate) == 0 {
			// TODO assure no invalid nets are created: offspring.Mutate()
			// offspring.Mutate()
		}

		e.Organisms = append(e.Organisms, NewOrganismFromNetwork(offspring.Decode(), e))
	}
}

func (e *Environment) RandomizeOrganisms() {
	coords := utils.GenerateUniqueCoordinates(e.MaxPop, e.Size, e.Size)

	for i, org := range e.Organisms {
		org.X = coords[i][0]
		org.Y = coords[i][1]
	}
}

// ApplySelection returns the number of organisms the selection function
// has applied to. If an organism was selected, it is removed from the
// environment.
func (e *Environment) ApplySelection() (deaths int) {
	survivors := e.Organisms[:0]
	for _, org := range e.Organisms {
		if e.selectionFn(org) {
			deaths++
		} else {
			survivors = append(survivors, org)
		}
	}
	e.Organisms = survivors
	return
}

func (e *Environment) Run() {
	for {
		if e.Paused {
			time.Sleep(100 * time.Millisecond)
			continue
		} else if e.MSPT > 0 {
			delta := time.Now().Sub(e.lastTick)
			if delta.Milliseconds() < e.MSPT {
				time.Sleep(time.Duration(e.MSPT-delta.Milliseconds()) * time.Millisecond)
			}
		}

		e.CurrGenAge++
		if e.CurrGenAge >= e.MaxGenAge {
			e.CurrGen++
			e.CurrGenAge = 0

			e.Survivors = e.MaxPop - e.ApplySelection()
		}

		if len(e.Organisms) < e.MaxPop {
			e.GenerateOffspring(800)
			e.RandomizeOrganisms()
		}

		for _, org := range e.Organisms {
			org.Compute()
		}

		e.lastTick = time.Now()
	}
}

func (e *Environment) Draw(screen *ebiten.Image) {
	size := float32(screen.Bounds().Max.X)
	cellSize := size / float32(e.Size)
	radius := float32(cellSize) / 2
	orgRadius := radius * 0.8

	vector.DrawFilledRect(screen, 0, 0,
		size, size,
		color.RGBA{R: 255, G: 255, B: 255}, false)

	for _, org := range e.Organisms {
		vector.DrawFilledCircle(
			screen,
			float32(org.X)*cellSize+radius,
			float32(org.Y)*cellSize+radius,
			orgRadius,
			org.Color,
			false)
	}
}
