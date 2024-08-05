package environment

import (
	"evolution/internal/environment/neural"
	"image/color"
)

type Organism struct {
	net        *neural.Network
	EncodedNet *neural.EncodedNet
	Color      color.RGBA

	X, Y   int
	Env    *Environment
	Facing neural.Orientation
}

func NewOrganism(nSynapses, maxInternalNeurons int, env *Environment) *Organism {
	net := neural.NewNeuralNet(nSynapses, maxInternalNeurons)
	encoded := net.Encode()
	return &Organism{
		net:        net,
		EncodedNet: encoded,
		Color:      encoded.Color(),
		Env:        env,
	}
}

func NewOrganismFromNetwork(net *neural.Network, env *Environment) *Organism {
	encoded := net.Encode()
	return &Organism{
		net:        net,
		EncodedNet: encoded,
		Color:      encoded.Color(),
		Env:        env,
	}
}

func (o *Organism) Compute() {
	o.net.Compute(o)
}

func (o *Organism) Move(d neural.Orientation) {
	newX, newY := o.X, o.Y
	switch d {
	case neural.North:
		if o.Y > 0 {
			newY--
		}
	case neural.East:
		if o.X < o.Env.Size-1 {
			newX++
		}
	case neural.South:
		if o.Y < o.Env.Size-1 {
			newY++
		}
	case neural.West:
		if o.X > 0 {
			newX--
		}
	}

	if !o.Env.IsOccupied(newX, newY) {
		o.X, o.Y = newX, newY
		o.Facing = d
	}
}

func (o *Organism) MoveDir(d neural.Direction) {
	o.Move(d.ToOrientation(o.Facing))
}
