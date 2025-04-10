package util

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	MplusNormalFace   *text.GoTextFace
	MplusNormalFaceSm *text.GoTextFace
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}

	MplusNormalFace = &text.GoTextFace{
		Source: s,
		Size:   24,
	}

	MplusNormalFaceSm = &text.GoTextFace{
		Source: s,
		Size:   16,
	}
}
