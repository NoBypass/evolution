package neural

import (
	"math"
	"math/rand"
)

type Kind int

type Type int

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

func (n *Neuron) Compute() float32 {
	parentSolutions := make([]float32, 0, len(n.Incoming))

	for _, synapse := range n.Incoming {
		parentSolutions = append(parentSolutions, synapse.From.Compute())
	}

	switch n.Type {
	case Sensory:
		return n.sensorData()
	case Internal:
		return n.computeInternal(parentSolutions)
	case Action:
		return n.computeAction(parentSolutions)
	default:
		panic("unknown neuron type")
	}
}

func (n *Neuron) sensorData() float32 {
	if n.Type != Sensory {
		panic("expected sensory neuron")
	}

	// TODO
	return rand.Float32()*2 - 1
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
