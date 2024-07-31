package environment

import "evolution/internal/environment/neural"

func (o *Organism) PopulationDensity() float32 {
	var surrounding int
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			if o.Env.IsOccupied(o.X+x, o.Y+y) {
				surrounding++
			}
		}
	}
	return float32(surrounding) / 8
}

func (o *Organism) BlockageFront() float32 {
	x, y := o.X, o.Y

	switch o.Facing {
	case neural.North:
		y--
	case neural.East:
		x++
	case neural.South:
		y++
	case neural.West:
		x--
	}

	if o.Env.IsOccupied(x, y) {
		return 1
	}
	return 0
}

func (o *Organism) BlockageLeftRight() float32 {
	xLeft, xRight, yLeft, yRight := o.X, o.X, o.Y, o.Y

	switch o.Facing {
	case neural.North:
		xLeft--
		xRight++
	case neural.East:
		yLeft--
		yRight++
	case neural.South:
		xLeft++
		xRight--
	case neural.West:
		yLeft++
		yRight--
	}

	if o.Env.IsOccupied(xLeft, yLeft) {
		return 1
	} else if o.Env.IsOccupied(xRight, yRight) {
		return -1
	}
	return 0
}

func (o *Organism) EastWestBorderDistance() float32 {
	return 2/float32(o.Env.Size)*float32(o.X) - 1
}

func (o *Organism) NorthSouthBorderDistance() float32 {
	return 2/float32(o.Env.Size)*float32(o.Y) - 1
}
