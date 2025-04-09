package game

import (
	"evolution/internal/environment"
	"evolution/internal/game/ui"
	"evolution/internal/graph"
	"fmt"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	HeaderOffset   = 48
	AnalysisOffset = 660
)

type Game struct {
	WindowSize int
	MainEnv    *environment.Environment

	currentEnv environment.Environment

	viewManager   *ui.ViewManager
	graphRenderer *graph.Renderer
}

func NewGame(size int, env *environment.Environment) *Game {
	g := Game{
		WindowSize: size,
		MainEnv:    env,
	}

	g.viewManager = ui.NewViewManager(size, ui.Handlers{
		PauseButton:     env.PauseHandler,
		MsptSubmit:      env.MSPTSubmitHandler,
		SimulationClick: g.orgSelectHandler,
	})
	g.graphRenderer = graph.NewRenderer(g.viewManager.GraphImg)

	go env.Run()
	go g.graphRenderer.Run()

	return &g
}

func (g *Game) Size() (int, int) {
	return g.WindowSize + AnalysisOffset, g.WindowSize + HeaderOffset
}

func (g *Game) Update() error {
	g.viewManager.UI.Update()

	g.currentEnv = *g.MainEnv

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.viewManager.GenerationLabel.Label = fmt.Sprintf("Gen: %d", g.currentEnv.CurrGen)
	g.viewManager.SurvivorsLabel.Label = fmt.Sprintf("Survivors: %d (%d%%)", g.currentEnv.Survivors, int(100/float32(g.currentEnv.MaxPop)*float32(g.currentEnv.Survivors)))

	g.currentEnv.Draw(g.viewManager.SimulationImg)
	g.graphRenderer.Draw(g.viewManager.GraphImg)

	g.viewManager.UI.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.Size()
}

func (g *Game) orgSelectHandler(args *widget.WidgetMouseButtonReleasedEventArgs) {
	if args.Inside {
		org := g.currentEnv.GetOrganismAt(image.Pt(args.OffsetX, args.OffsetY), g.WindowSize)
		if org != nil {
			g.graphRenderer.OrgCh <- org
		}
	}
}
