package util

import (
	"github.com/chewxy/math32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Options struct {
	Color     color.Color
	AntiAlias bool
	Width     float32
}

type Option func(*Options)

var sqrt3 = math32.Sqrt(3)

func DrawLine(img *ebiten.Image, v1, v2 Vector2, options ...Option) {
	opts := applyOpts(options)
	vector.StrokeLine(img, v1.X, v1.Y, v2.X, v2.Y, opts.Width, opts.Color, opts.AntiAlias)
}

func DrawFilledCircle(img *ebiten.Image, at Vector2, radius float32, options ...Option) {
	opts := applyOpts(options)
	vector.DrawFilledCircle(img, at.X, at.Y, radius, opts.Color, opts.AntiAlias)
}

func DrawText(img *ebiten.Image, str string, at Vector2, face text.Face) {
	op := &text.DrawOptions{}
	op.ColorScale.Scale(0, 0, 0, 1)
	op.GeoM.Translate(at.As64())
	op.PrimaryAlign = text.AlignCenter
	op.SecondaryAlign = text.AlignCenter
	text.Draw(img, str, face, op)
}

func DrawFilledArrow(img *ebiten.Image, from, to Vector2, options ...Option) {
	opts := applyOpts(options)
	r, g, b, a := opts.Color.RGBA()

	distVec := to.Sub(from)
	distVecLen := distVec.Len()
	sideLen := distVecLen / sqrt3

	orthogonalRight := Vec2(distVec.Y, -distVec.X).Normalize().Mul(sideLen)
	orthogonalLeft := Vec2(-distVec.Y, distVec.X).Normalize().Mul(sideLen)

	vertices := []ebiten.Vertex{
		to.AsVertex(),
		from.Add(orthogonalRight).AsVertex(),
		from.Add(orthogonalLeft).AsVertex(),
	}

	for i := range vertices {
		vertex := &vertices[i]
		vertex.ColorR = float32(r) / 0xffff
		vertex.ColorG = float32(g) / 0xffff
		vertex.ColorB = float32(b) / 0xffff
		vertex.ColorA = float32(a) / 0xffff
	}

	whiteImage := ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)

	img.DrawTriangles(vertices, []uint16{0, 1, 2}, whiteImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: opts.AntiAlias,
	})
}

func applyOpts(opts []Option) *Options {
	opt := &Options{
		Color:     color.Black,
		AntiAlias: true,
		Width:     1,
	}

	for _, fn := range opts {
		fn(opt)
	}

	return opt
}

func WithColor(c color.Color) Option {
	return func(o *Options) {
		o.Color = c
	}
}

func WithAntiAlias(aa bool) Option {
	return func(o *Options) {
		o.AntiAlias = aa
	}
}

func WithWidth(w float32) Option {
	return func(o *Options) {
		o.Width = w
	}
}
