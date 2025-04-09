package environment

import (
	"encoding/base64"
	"evolution/internal/environment/neural"
	"image/color"
)

type Organism struct {
	Net        *neural.Network
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
		Net:        net,
		EncodedNet: encoded,
		Color:      encoded.Color(),
		Env:        env,
	}
}

func NewOrganismFromNetwork(net *neural.Network, env *Environment) *Organism {
	encoded := net.Encode()
	return &Organism{
		Net:        net,
		EncodedNet: encoded,
		Color:      encoded.Color(),
		Env:        env,
	}
}

func NewOrganismFromEncodedNet(encodedNet string) *Organism {
	data, err := base64.RawURLEncoding.DecodeString(encodedNet)
	if err != nil {
		panic(err)
	}

	en := make(neural.EncodedNet, len(data)/6)
	for i := range en {
		copy(en[i][:], data[i*6:(i+1)*6])
	}

	return &Organism{
		Net:        en.Decode(),
		EncodedNet: &en,
		Color:      en.Color(),
	}
}

func (o *Organism) Compute() {
	o.Net.Compute(o)
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

func (o *Organism) String() string {
	en := o.EncodedNet
	buf := make([]byte, len(*en)*6)
	for i, es := range *en {
		copy(buf[i*6:], es[:])
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}
