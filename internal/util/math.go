package util

import (
	"fmt"
	"github.com/chewxy/math32"
	"github.com/hajimehoshi/ebiten/v2"
)

type Vector2 struct {
	X, Y float32
}

func Vec2(X, Y float32) Vector2 {
	return Vector2{X, Y}
}

func (v Vector2) Add(other Vector2) Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vector2) Sub(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vector2) Div(n float32) Vector2 {
	v.X /= n
	v.Y /= n
	return v
}

func (v Vector2) Mul(n float32) Vector2 {
	v.X *= n
	v.Y *= n
	return v
}

func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) EProd(other Vector2) Vector2 {
	return Vec2(v.X*other.X, v.Y*other.Y)
}

func (v Vector2) Len() float32 {
	return math32.Hypot(v.X, v.Y)
}

func (v Vector2) Distance(other Vector2) float32 {
	return math32.Hypot(other.X, other.Y)
}

func (v Vector2) Normalize() Vector2 {
	return v.Div(v.Len())
}

func (v Vector2) As64() (float64, float64) {
	return float64(v.X), float64(v.Y)
}

func (v Vector2) AsVertex() ebiten.Vertex {
	return ebiten.Vertex{
		DstX: v.X,
		DstY: v.Y,
	}
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
