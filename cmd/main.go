package main

import (
	"evolution/internal/environment"
	"evolution/internal/game"
	"evolution/internal/web"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const WindowSize = 640

func main() {
	go web.RunServer()

	env := environment.New(100, 1000, 128, func(org *environment.Organism) bool {
		return org.EastWestBorderDistance() < 0
	})

	g := game.NewGame(WindowSize, env)
	ebiten.SetWindowSize(g.Size())
	ebiten.SetWindowTitle("Evolution")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
