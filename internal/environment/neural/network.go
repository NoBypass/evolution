package neural

import (
	"fmt"
	"math/rand"
)

type Network struct {
	// neurons is a slice of sensory neurons. Each neuron
	// has links to internal and/or action neurons.
	neurons []*Neuron
}

func NewNeuralNet(nSynapses, nInternal int) *Network {
	synapses := make([]*Synapse, 0, nSynapses)
	neurons := make(map[Kind]*Neuron)

	for range nSynapses {
		from := randomSensoryOrInternalNeuron(nInternal)
		var to *Neuron
		if from.Type == Internal {
			to = randomInternalOrActionNeuron(0)
		} else {
			to = randomInternalOrActionNeuron(nInternal)
		}

		if existing, exists := neurons[from.ID]; exists {
			from = existing
		} else {
			neurons[from.ID] = from
		}

		if existing, exists := neurons[to.ID]; exists {
			to = existing
		} else {
			neurons[to.ID] = to
		}

		synapses = append(synapses, &Synapse{
			Weight: rand.Float32()*8 - 4,
			From:   from,
			To:     to,
		})
	}

	nn := &Network{
		neurons: make([]*Neuron, 0, len(neurons)),
	}

	for _, synapse := range synapses {
		synapse.From.Outgoing = synapse
		synapse.To.Incoming = append(synapse.To.Incoming, synapse)
	}

	for _, neuron := range neurons {
		if neuron.Type == Action {
			nn.neurons = append(nn.neurons, neuron)
		}
	}

	return nn
}

func (n *Network) Compute(exec IAction) {
	highest := float32(-1)
	action := new(Neuron)

	for _, actionNeuron := range n.neurons {
		num := actionNeuron.Compute()
		if num > highest {
			action = actionNeuron
			highest = num
		}
	}

	switch action.ID {
	case MoveForward:
		exec.MoveDir(Front)
	case MoveBackward:
		exec.MoveDir(Back)
	case MoveRandom:
		exec.Move(Orientation(rand.Intn(4)))
	case MoveLeftRight:
		if highest > 0 {
			exec.MoveDir(Left)
		} else {
			exec.MoveDir(Right)
		}
	case MoveEastWest:
		if highest > 0 {
			exec.Move(East)
		} else {
			exec.Move(West)
		}
	case MoveNorthSouth:
		if highest > 0 {
			exec.Move(North)
		} else {
			exec.Move(South)
		}
	default:
		panic(fmt.Sprintf("unexpected action %d", action.ID))
	}
}
