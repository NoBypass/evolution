package main

import (
	"evolution/internal/environment"
	"evolution/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const WindowSize = 640

func main() {
	g := game.Game{
		// TODO: not hardcode values
		WindowSize: WindowSize,
		Env: environment.New(96, 1000, 128, func(org *environment.Organism) bool {
			return org.EastWestBorderDistance() < 0
		}),
	}
	ebiten.SetWindowSize(g.Size())
	ebiten.SetWindowTitle("Evolution")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
