package neural

import (
	"math"
	"math/rand"
)

type Kind int8

type Type int8

type Neuron struct {
	ID       Kind
	Type     Type
	Incoming []*Synapse
	Outgoing *Synapse
}

type Synapse struct {
	Weight   float32
	From, To *Neuron
}

type SensorEquipped interface {
	// PopulationDensity a number between 0 and 1 representing the
	// fraction of surrounding cells that are occupied.
	PopulationDensity() float32
	// BlockageFront returns 1 if the cell in front of the organism
	// is occupied, 0 otherwise.
	BlockageFront() float32
	// BlockageLeftRight returns 1 if the cell to the left is occupied,
	// -1 if the cell to the right is occupied, 0 otherwise.
	BlockageLeftRight() float32
	// EastWestBorderDistance returns a positive number for the fraction
	// of the distance from the organism to the east border and a negative
	// number for the distance to the west border.
	EastWestBorderDistance() float32
	// NorthSouthBorderDistance returns a positive number for the fraction
	// of the distance from the organism to the south border and a negative
	// number for the distance to the north border.
	NorthSouthBorderDistance() float32
}

func (n *Neuron) Compute(se SensorEquipped) float32 {
	parentSolutions := make([]float32, 0, len(n.Incoming))

	for _, synapse := range n.Incoming {
		parentSolutions = append(parentSolutions, synapse.From.Compute(se))
	}

	switch n.Type {
	case Sensory:
		return n.sensorData(se)
	case Internal:
		return n.computeInternal(parentSolutions)
	case Action:
		return n.computeAction(parentSolutions)
	default:
		panic("unknown neuron type")
	}
}

func (n *Neuron) sensorData(se SensorEquipped) float32 {
	if n.Type != Sensory {
		panic("expected sensory neuron")
	}

	switch n.ID {
	case Random:
		return rand.Float32()*2 - 1
	case PopDensity:
		return se.PopulationDensity()
	case BlockageFront:
		return se.BlockageFront()
	case BlockageLeftRight:
		return se.BlockageLeftRight()
	case EastWestBorderDistance:
		return se.EastWestBorderDistance()
	case NorthSouthBorderDistance:
		return se.NorthSouthBorderDistance()
	default:
		panic("unexpected sensory neuron")
	}
}

func (n *Neuron) computeInternal(solutions []float32) float32 {
	if n.Type != Internal {
		panic("expected internal neuron")
	}

	var sum float64
	for _, s := range solutions {
		sum += float64(s)
	}

	return float32(math.Tanh(sum))
}

func (n *Neuron) computeAction(solutions []float32) float32 {
	if n.Type != Action {
		panic("expected action neuron")
	}

	var sum float64
	for _, s := range solutions {
		sum += float64(s)
	}

	return float32(math.Tanh(sum))
}
