package game

import (
	"evolution/internal/environment"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	HeaderOffset   = 48
	AnalysisOffset = 660
)

type Game struct {
	WindowSize int
	MainEnv    *environment.Environment

	currentEnv environment.Environment

	clickedLastFrame bool
	viewManager      *viewManager
}

func NewGame(size int, env *environment.Environment) *Game {
	g := Game{
		WindowSize: size,
		MainEnv:    env,

		viewManager: newViewManager(size, handlers{
			pauseButton: env.PauseHandler,
			msptSubmit:  env.MSPTSubmitHandler,
		}),
	}

	go env.Run()

	return &g
}

func (g *Game) Size() (int, int) {
	return g.WindowSize + AnalysisOffset, g.WindowSize + HeaderOffset
}

func (g *Game) Update() error {
	g.viewManager.ui.Update()

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
	g.viewManager.generationLabel.Label = fmt.Sprintf("Gen: %d", g.currentEnv.CurrGen)
	g.viewManager.survivorsLabel.Label = fmt.Sprintf("Survivors: %d (%d%%)", g.currentEnv.Survivors, int(100/float32(g.currentEnv.MaxPop)*float32(g.currentEnv.Survivors)))

	g.currentEnv.Draw(g.viewManager.simulationImg)

	g.viewManager.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Size()
}
