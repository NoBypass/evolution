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
		WindowSize: WindowSize,
		// TODO: not hardcode nSynapses and maxInternalNeurons
		Env: environment.New(128, 1000),
	}
	ebiten.SetWindowSize(WindowSize, WindowSize)
	ebiten.SetWindowTitle("Evolution")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
