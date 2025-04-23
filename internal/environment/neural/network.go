package neural

import (
	"fmt"
	"github.com/chewxy/math32"
	"math/rand"
	"slices"
)

type Network struct {
	Neurons  []*Neuron
	Synapses map[*Synapse]float32
}

func NewNeuralNet(synapseCount, internalCount int) *Network {
	sample := make(map[Kind]*Neuron, _totalNeurons)
	inDegree := make(map[*Neuron]int, _totalNeurons)

	for i := range _totalNeurons + internalCount {
		k := Kind(i - internalCount)
		sample[k] = &Neuron{ID: k, Type: k.asType()}
	}

	net := &Network{
		Neurons:  make([]*Neuron, 0),
		Synapses: make(map[*Synapse]float32),
	}

	dagSlice := generateDAGSlice(internalCount)

	for range synapseCount {
		fromK := Kind(rand.Intn(_totalSensoryNeurons+internalCount) - internalCount)

		var selection []Kind
		if fromK.asType() == Sensory {
			selection = dagSlice[_totalSensoryNeurons:]
		} else {
			selection = dagSlice[slices.Index(dagSlice, fromK)+1:]
		}

		toK := selection[rand.Intn(len(selection))]

		inDegree[sample[fromK]], inDegree[sample[toK]] = 0, 0

		net.Synapses[&Synapse{
			Weight: rand.Float32()*8 - 4,
			From:   sample[fromK],
			To:     sample[toK],
		}] = 0
	}

	for s := range net.Synapses {
		inDegree[s.To]++
	}

	queue := make([]*Neuron, 0)

	for id, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, id)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		net.Neurons = append(net.Neurons, current)

		for s := range net.Synapses {
			if s.From == current {
				inDegree[s.To]--
				if inDegree[s.To] == 0 {
					queue = append(queue, s.To)
				}
			}
		}
	}

	if len(net.Neurons) != len(inDegree) {
		panic("neural network contains a cycle, topological sort not possible")
	}

	return net
}

func generateDAGSlice(internalCount int) []Kind {
	s := make([]Kind, _totalNeurons+internalCount)
	for i := range s {
		s[i] = Kind(i - internalCount)
	}

	at := _totalSensoryNeurons + internalCount
	result := make([]Kind, len(s))
	copy(result, s[internalCount:at])
	copy(result[len(s[internalCount:at]):], s[:internalCount])
	copy(result[len(s[internalCount:at])+len(s[:internalCount]):], s[at:])
	return result
}

type MovableSE interface {
	SensorEquipped
	Movable
}

func (n *Network) Compute(exec MovableSE) {
	for _, neuron := range n.Neurons {
		if neuron.Type == Sensory {
			neuron.Value = neuron.SensorData(exec)
		} else {
			for synapse, v := range n.Synapses {
				if synapse.To == neuron {
					neuron.Value += v
				}
			}

			neuron.Value = math32.Tanh(neuron.Value)
		}

		for synapse := range n.Synapses {
			if synapse.From == neuron {
				n.Synapses[synapse] = synapse.Weight * neuron.Value
			}
		}
	}

	var highest *Neuron
	for _, neuron := range n.Neurons {
		if neuron.Type == Action && (highest == nil || math32.Abs(neuron.Value) > math32.Abs(highest.Value)) {
			highest = neuron
		}
	}

	if highest == nil {
		panic("no highest action neuron could be found")
	}

	switch highest.ID {
	case MoveForward:
		exec.MoveDir(Front)
	case MoveBackward:
		exec.MoveDir(Back)
	case MoveRandom:
		exec.Move(Orientation(rand.Intn(4)))
	case MoveLeftRight:
		if highest.Value > 0 {
			exec.MoveDir(Left)
		} else {
			exec.MoveDir(Right)
		}
	case MoveEastWest:
		if highest.Value > 0 {
			exec.Move(East)
		} else {
			exec.Move(West)
		}
	case MoveNorthSouth:
		if highest.Value > 0 {
			exec.Move(North)
		} else {
			exec.Move(South)
		}
	default:
		panic(fmt.Sprintf("unexpected action %d", highest.ID))
	}
}
