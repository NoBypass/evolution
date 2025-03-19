package game

import (
	"evolution/internal/command"
	"evolution/internal/environment"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"strconv"
)

const (
	HeaderOffset = 48
)

type Game struct {
	WindowSize int
	MainEnv    *environment.Environment

	UpdateRequester chan<- struct{}
	UpdateReceiver  <-chan environment.Environment
	currentEnv      *environment.Environment

	cmdHandler *command.Handler
}

func NewGame(size int, env *environment.Environment) *Game {
	g := Game{
		WindowSize: size,
		MainEnv:    env,

		UpdateRequester: env.InitRequester(),
		UpdateReceiver:  env.InitReceiver(),

		cmdHandler: command.NewHandler(),
	}

	go env.Run()
	go g.cmdHandler.Run()

	return &g
}

func (g *Game) Size() (int, int) {
	return g.WindowSize, g.WindowSize + HeaderOffset
}

func (g *Game) Update() error {
	select {
	case cmd := <-g.cmdHandler.Ch:
		switch cmd.Command {
		case "mspt":
			mspt, err := strconv.ParseInt(cmd.Value, 10, 64)
			if err != nil {
				fmt.Printf("Invalid mspt value %s\n", cmd.Value)
			}

			g.MainEnv.MSPT = mspt
		}
	default:
	}

	g.UpdateRequester <- struct{}{}
	env := <-g.UpdateReceiver
	g.currentEnv = &env
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
