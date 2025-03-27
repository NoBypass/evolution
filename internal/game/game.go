package game

import (
	"evolution/internal/command"
	"evolution/internal/environment"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	HeaderOffset = 48
)

type Game struct {
	WindowSize int
	MainEnv    *environment.Environment

	currentEnv environment.Environment
	cmdHandler *command.Handler

	clickedLastFrame bool
}

func NewGame(size int, env *environment.Environment) *Game {
	g := Game{
		WindowSize: size,
		MainEnv:    env,

		cmdHandler: command.NewHandler(),
	}

	go env.Run()
	go g.cmdHandler.Run(g.MainEnv)

	return &g
}

func (g *Game) Size() (int, int) {
	return g.WindowSize, g.WindowSize + HeaderOffset
}

func (g *Game) Update() error {
	g.currentEnv = *g.MainEnv

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.clickedLastFrame {
			g.clickedLastFrame = true
			x, y := ebiten.CursorPosition()
			org := g.getOrganismAtCoordinates(float32(x), float32(y))
			if org != nil {
				fmt.Printf("Clicked on organism: %s\n", org)
			}
		}
	} else {
		g.clickedLastFrame = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)

	g.drawText(screen, fmt.Sprintf("Gen: %d", g.currentEnv.CurrGen), 8, 8)
	g.drawText(screen, fmt.Sprintf("Survivors: %d (%d%%)", g.currentEnv.Survivors, int(100/float32(g.currentEnv.MaxPop)*float32(g.currentEnv.Survivors))), 160, 8)

	g.drawEnvironment(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Size()
}
