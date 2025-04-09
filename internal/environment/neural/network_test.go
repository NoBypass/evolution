package neural

import (
	"slices"
	"testing"
)

func TestNewNeuralNet(t *testing.T) {
	neurons := make(map[*Neuron]struct{})

	nn := NewNeuralNet(20, 3)
	for _, net := range nn.ActionNeurons {
		deconstructNeuron(neurons, net)
	}

	t.Run("test if kind matches type", func(t *testing.T) {
		possibleSensory := []Kind{Random, PopDensity, BlockageFront, BlockageLeftRight, EastWestBorderDistance, NorthSouthBorderDistance}
		possibleAction := []Kind{MoveForward, MoveBackward, MoveRandom, MoveLeftRight, MoveEastWest, MoveNorthSouth}
		possibleInternal := []Kind{-1, -2, -3}

		for neuron := range neurons {
			switch neuron.Type {
			case Sensory:
				if !slices.Contains(possibleSensory, neuron.ID) {
					t.Errorf("neural %v is sensory but has kind %d", neuron, neuron.ID)
				}
			case Internal:
				if !slices.Contains(possibleInternal, neuron.ID) {
					t.Errorf("neural %v is internal but has kind %d", neuron, neuron.ID)
				}
			case Action:
				if !slices.Contains(possibleAction, neuron.ID) {
					t.Errorf("neural %v is action but has kind %d", neuron, neuron.ID)
				}
			}
		}
	})
	t.Run("test if all heads are action ActionNeurons", func(t *testing.T) {
		for _, actionNeurons := range nn.ActionNeurons {
			if actionNeurons.Type != Action {
				t.Errorf("neural %v is not an action neuron", actionNeurons)
			}
		}
	})
}

func deconstructNeuron(neurons map[*Neuron]struct{}, neuron *Neuron) {
	neurons[neuron] = struct{}{}

	if neuron.Outgoing != nil {
		deconstructNeuron(neurons, neuron.Outgoing.To)
	}
}
