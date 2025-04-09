package graph

import (
	"evolution/internal/environment"
	"evolution/internal/environment/neural"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	NodeRadius = 25
)

type Renderer struct {
	OrgCh chan *environment.Organism

	img   *ebiten.Image
	graph *Graph
}

func NewRenderer(img *ebiten.Image) *Renderer {
	return &Renderer{
		img:   img,
		graph: &Graph{},

		OrgCh: make(chan *environment.Organism),
	}
}

func (r *Renderer) Run() {
	for org := range r.OrgCh {
		r.graph = new(Graph)
		m := make(map[*neural.Neuron]*Node)
		for _, n := range org.Net.ActionNeurons {
			traverse(n, m)
		}

		for _, n := range m {
			r.graph.Nodes = append(r.graph.Nodes, n)
		}

	ticker:
		for {
			select {
			case org = <-r.OrgCh:
				break ticker
			default:
				r.graph.tick()
			}
		}
	}
}

func traverse(neuron *neural.Neuron, m map[*neural.Neuron]*Node) *Node {
	node := &Node{
		X: Width / 2,
		Y: Height / 2,
	}
	for _, s := range neuron.Incoming {
		var to *Node
		if n, ok := m[neuron]; ok {
			to = n
		} else {
			to = traverse(s.From, m)
		}

		node.Edges = append(node.Edges, &Edge{
			Weight: float64(s.Weight),
			To:     to,
		})
	}

	m[neuron] = node
	return node
}

func (r *Renderer) Draw(img *ebiten.Image) {
	img.Clear()
	for _, n := range r.graph.Nodes {
		vector.DrawFilledCircle(
			img,
			float32(n.X),
			float32(n.Y),
			NodeRadius,
			color.Black,
			false)
		for _, e := range n.Edges {
			vector.StrokeLine(
				img,
				float32(n.X),
				float32(n.Y),
				float32(e.To.X),
				float32(e.To.Y),
				2,
				color.Black,
				false)
		}
	}
}
