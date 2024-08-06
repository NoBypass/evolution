package game

import (
	"evolution/internal/environment"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	HeaderOffset = 48
)

type Game struct {
	WindowSize int
	Env        *environment.Environment

	sleep     *time.Timer
	survivors int
}

func (g *Game) Size() (int, int) {
	return g.WindowSize, g.WindowSize + HeaderOffset
}

func (g *Game) Update() error {
	if g.sleep != nil {
		select {
		case <-g.sleep.C:
			g.sleep = nil
		default:
			return nil
		}
	}

	g.Env.CurrGenAge++
	if g.Env.CurrGenAge >= g.Env.MaxGenAge {
		g.Env.CurrGen++
		g.Env.CurrGenAge = 0

		g.survivors = g.Env.MaxPop - g.Env.ApplySelection()
		g.sleep = time.NewTimer(75 * time.Millisecond)
		return nil
	}
	if len(g.Env.Organisms) < g.Env.MaxPop {
		g.Env.GenerateOffspring(800)
		g.Env.RandomizeOrganisms()
	}

	for _, org := range g.Env.Organisms {
		org.Compute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)

	g.drawText(screen, fmt.Sprintf("Gen: %d", g.Env.CurrGen), 8, 8)
	g.drawText(screen, fmt.Sprintf("Survivors: %d (%d%%)", g.survivors, int(100/float32(g.Env.MaxPop)*float32(g.survivors))), 160, 8)

	g.drawEnvironment(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Size()
}
