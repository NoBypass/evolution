package neural

import (
	"slices"
	"testing"
)

func TestNeuronGenerators(t *testing.T) {
	t.Run("test random internal or action neural", func(t *testing.T) {
		possibleNeurons := []int{-3, -2, -1, 9, 10, 11, 12, 13, 14}
		got := make(map[int]int)
		for range 1000 {
			neuron := randomInternalOrActionNeuron(3)
			got[int(neuron.ID)]++
		}

		for neuron, count := range got {
			if !slices.Contains(possibleNeurons, neuron) {
				t.Errorf("unexpected neural %d generated %d times", neuron, count)
			} else {
				t.Logf("neural %d generated %d times", neuron, count)
			}
		}

		for _, neuron := range possibleNeurons {
			if _, ok := got[neuron]; !ok {
				t.Errorf("neural %d was never generated", neuron)
			}
		}
	})
	t.Run("test random sensory or internal neural", func(t *testing.T) {
		possibleNeurons := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7}
		got := make(map[int]int)
		for range 1000 {
			neuron := randomSensoryOrInternalNeuron(3)
			got[int(neuron.ID)]++
		}

		for neuron, count := range got {
			if !slices.Contains(possibleNeurons, neuron) {
				t.Errorf("unexpected neural %d generated %d times", neuron, count)
			} else {
				t.Logf("neural %d generated %d times", neuron, count)
			}
		}

		for _, neuron := range possibleNeurons {
			if _, ok := got[neuron]; !ok {
				t.Errorf("neural %d was never generated", neuron)
			}
		}
	})
}
