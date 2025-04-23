package neural

const (
	Random Kind = iota
	PopDensity
	BlockageFront
	BlockageLeftRight
	EastWestBorderDistance
	NorthSouthBorderDistance

	_totalSensoryNeurons int = iota

	MoveForward Kind = iota - 1
	MoveBackward
	MoveRandom
	MoveLeftRight
	MoveEastWest
	MoveNorthSouth

	_totalNeurons int = iota - 1
)

const (
	Sensory Type = iota
	Internal
	Action
)

func generateNeuronByKind(kind Kind) *Neuron {
	t := Action
	if kind < 0 {
		t = Internal
	} else if int(kind) < _totalSensoryNeurons {
		t = Sensory
	}

	return &Neuron{
		ID:   kind,
		Type: t,
	}
}

func (k Kind) asType() Type {
	if k < 0 {
		return Internal
	} else if int(k) < _totalSensoryNeurons {
		return Sensory
	} else {
		return Action
	}
}
