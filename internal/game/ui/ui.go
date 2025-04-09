package ui

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
	"image/png"
)

var (
	mplusNormalFace *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}

	mplusNormalFace = &text.GoTextFace{
		Source: s,
		Size:   24,
	}
}

func newPlainImage(size int) *ebiten.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	ebitenImg, _, err := ebitenutil.NewImageFromReader(&buf)
	if err != nil {
		panic(err)
	}

	return ebitenImg
}
