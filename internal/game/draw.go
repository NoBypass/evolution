package game

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

var (
	mplusFaceSource *text.GoTextFaceSource
	mplusNormalFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	mplusNormalFace = &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	width, height := g.Size()
	vector.DrawFilledRect(screen, 0, 0,
		float32(width), float32(height),
		color.RGBA{R: 255, G: 255, B: 255}, false)
}

func (g *Game) drawText(screen *ebiten.Image, str string, x, y int) {
	op := &text.DrawOptions{}
	op.ColorScale.Scale(0, 0, 0, 1)
	op.GeoM.Translate(float64(x), float64(y))
	text.Draw(screen, str, mplusNormalFace, op)
}

func (g *Game) drawEnvironment(screen *ebiten.Image) {
	cellSize := float32(g.WindowSize) / float32(g.Env.Size)
	radius := float32(cellSize) / 2
	orgRadius := radius * 0.8

	for _, org := range g.Env.Organisms {
		vector.DrawFilledCircle(
			screen,
			float32(org.X)*cellSize+radius,
			float32(org.Y)*cellSize+radius+HeaderOffset,
			orgRadius,
			org.Color,
			false)
	}
}
