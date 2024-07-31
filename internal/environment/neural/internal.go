package neural

import "math/rand"

const (
	Random Kind = iota
	PopDensity
	BlockageFront
	BlockageLeftRight
	EastWestBorderDistance
	NorthSouthBorderDistance

	_totalSensoryNeurons

	MoveForward
	MoveBackward
	MoveRandom
	MoveLeftRight
	MoveEastWest
	MoveNorthSouth

	_totalNeurons
)

const (
	Sensory Type = iota
	Internal
	Action
)

func randomSensoryOrInternalNeuron(nInternal int) *Neuron {
	kind := Kind(rand.Intn(int(_totalSensoryNeurons)+nInternal) - nInternal)
	t := Sensory
	if kind < 0 {
		t = Internal
	}

	return &Neuron{
		ID:       kind,
		Type:     t,
		Incoming: make([]*Synapse, 0),
	}
}

func randomInternalOrActionNeuron(nInternal int) *Neuron {
	kind := Kind(rand.Intn(int(_totalNeurons)-(1+int(_totalSensoryNeurons))+nInternal) + 1 + int(_totalSensoryNeurons))
	t := Action

	if kind >= _totalNeurons {
		kind = 0 - kind + _totalNeurons - 1
		t = Internal
	}

	return &Neuron{
		ID:       kind,
		Type:     t,
		Incoming: make([]*Synapse, 0),
	}
}
