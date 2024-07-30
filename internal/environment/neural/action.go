package neural

type Orientation int
type Direction int

const (
	North Orientation = iota
	East
	South
	West
)

const (
	Front Direction = iota
	Right
	Back
	Left
)

type IAction interface {
	Move(d Orientation)
	MoveDir(d Direction)
}

func (d Direction) ToOrientation(facing Orientation) Orientation {
	return ((facing) + Orientation(d)) % 4
}
