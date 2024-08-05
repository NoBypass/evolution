package game

import (
	"evolution/internal/environment"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Game struct {
	WindowSize int
	Env        *environment.Environment

	resetting bool
}

func (g *Game) Update() error {
	g.Env.CurrGenAge++
	if g.Env.CurrGenAge >= g.Env.MaxGenAge {
		g.Env.CurrGenAge = 0

		g.Env.ApplySelection()
		g.Env.GenerateOffspring(800)
		g.Env.RandomizeOrganisms()

		return nil
	}

	for _, org := range g.Env.Organisms {
		org.Compute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cellSize := float32(g.WindowSize) / float32(g.Env.Size)
	radius := float32(cellSize) / 2
	orgRadius := radius * 0.8

	vector.DrawFilledRect(screen, 0, 0,
		float32(g.WindowSize), float32(g.WindowSize),
		color.RGBA{R: 255, G: 255, B: 255}, false)

	for _, org := range g.Env.Organisms {
		vector.DrawFilledCircle(
			screen,
			float32(org.X)*cellSize+radius,
			float32(org.Y)*cellSize+radius,
			orgRadius,
			org.Color,
			false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.WindowSize, g.WindowSize
}
