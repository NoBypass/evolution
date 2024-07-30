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
}

func (g *Game) Update() error {
	for _, org := range g.Env.Organisms {
		org.Compute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cellSize := float32(g.WindowSize) / float32(g.Env.Size)
	radius := float32(cellSize) / 2
	vector.DrawFilledRect(screen, 0, 0,
		float32(g.WindowSize), float32(g.WindowSize),
		color.RGBA{R: 255, G: 255, B: 255}, false)

	for _, org := range g.Env.Organisms {
		vector.DrawFilledCircle(
			screen,
			float32(org.X)*cellSize+radius,
			float32(org.Y)*cellSize+radius,
			radius*0.8,
			color.RGBA{R: 255, A: 255},
			false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.WindowSize, g.WindowSize
}
