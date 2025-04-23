package neural

import (
	"encoding/binary"
	"image/color"
	"math"
	"math/rand"
)

type EncodedNet []EncodedSynapse
type EncodedSynapse [6]byte

func (s *Synapse) encode() EncodedSynapse {
	var b [6]byte
	binary.LittleEndian.PutUint32(b[:4], math.Float32bits(s.Weight))
	b[4] = byte(s.From.ID)
	b[5] = byte(s.To.ID)
	return b
}

func (es *EncodedSynapse) weight() float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32((*es)[:4]))
}

func (es *EncodedSynapse) fromKind() Kind {
	return Kind((*es)[4])
}

func (es *EncodedSynapse) toKind() Kind {
	return Kind((*es)[5])
}

func (en *EncodedNet) Color() color.RGBA {
	var r, g, b float32
	for _, es := range *en {
		r += float32(es.fromKind())
		g += float32(es.toKind())
		b += es.weight()
	}

	return color.RGBA{
		R: uint8(r / float32(len(*en))),
		G: uint8(g / float32(len(*en))),
		B: uint8(b / float32(len(*en))),
		A: 255,
	}
}

func (en *EncodedNet) Mutate() {
	(*en)[rand.Intn(len(*en))][rand.Intn(6)] = byte(rand.Intn(256))
}

func (n *Network) Encode() *EncodedNet {
	encoded := make([]EncodedSynapse, 0, len(n.Synapses))
	for synapse := range n.Synapses {
		encoded = append(encoded, synapse.encode())
	}
	return (*EncodedNet)(&encoded)
}

func OffspringOf(a, b *EncodedNet) *EncodedNet {
	offspring := make([]EncodedSynapse, 0, len(*a))
	for i := range *a {
		if i%2 == 0 {
			offspring = append(offspring, (*a)[i])
		} else {
			offspring = append(offspring, (*b)[i])
		}
	}
	return (*EncodedNet)(&offspring)
}

func (en *EncodedNet) Decode() *Network {
	neurons := make(map[Kind]*Neuron)
	synapses := make(map[*Synapse]float32, len(*en))

	for _, es := range *en {
		fromKind, toKind := es.fromKind(), es.toKind()
		from, to := neurons[fromKind], neurons[toKind]

		if from == nil {
			from = generateNeuronByKind(fromKind)
			neurons[fromKind] = from
		}

		if to == nil {
			to = generateNeuronByKind(toKind)
			neurons[toKind] = to
		}

		synapses[&Synapse{
			Weight: es.weight(),
			From:   from,
			To:     to,
		}] = 0
	}

	nn := &Network{
		Neurons:  make([]*Neuron, 0, len(neurons)),
		Synapses: synapses,
	}

	for _, neuron := range neurons {
		if neuron.Type == Action {
			nn.Neurons = append(nn.Neurons, neuron)
		}
	}

	return nn
}
