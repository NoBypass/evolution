package graph

import (
	"evolution/internal/environment"
	"evolution/internal/environment/neural"
	"evolution/internal/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	NodeRadius = 25
)

type Renderer struct {
	OrgCh chan *environment.Organism

	img   *ebiten.Image
	graph *Graph

	axisMultiplier util.Vector2
}

func NewRenderer(img *ebiten.Image) *Renderer {
	return &Renderer{
		img:   img,
		graph: &Graph{},

		OrgCh:          make(chan *environment.Organism),
		axisMultiplier: util.Vec2(float32(img.Bounds().Max.X)/graphWidth, float32(img.Bounds().Max.Y)/graphHeight),
	}
}

func (r *Renderer) Run() {
	org := <-r.OrgCh
	for {
		r.graph = new(Graph)
		nodes := make(map[*neural.Neuron]*Node)

		for _, syn := range org.Net.Synapses {
			to, ok := nodes[syn.To]
			if !ok {
				to = newNode(syn.To)
				nodes[syn.To] = to
			}

			edge := &Edge{
				Weight: syn.Weight,
				To:     to,
			}

			from, ok := nodes[syn.From]
			if !ok {
				from = newNode(syn.From)
				nodes[syn.From] = from
			}

			from.Edges = append(from.Edges, edge)
		}

		for _, node := range nodes {
			r.graph.Nodes = append(r.graph.Nodes, node)
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

func newNode(neuron *neural.Neuron) *Node {
	var clr color.Color = color.Black
	switch neuron.Type {
	case neural.Sensory:
		clr = color.RGBA{94, 233, 181, 255}
	case neural.Internal:
		clr = color.RGBA{212, 212, 212, 255}
	case neural.Action:
		clr = color.RGBA{255, 184, 106, 255}
	}

	return &Node{
		pos:   centerVec,
		Color: clr,
		Text:  neuron.ID.String(),
		Edges: make([]*Edge, 0),
	}
}

func (r *Renderer) Draw(img *ebiten.Image) {
	img.Clear()
	for _, n := range r.graph.Nodes {
		vNode := n.pos.EProd(r.axisMultiplier)

		util.DrawFilledCircle(img, vNode, NodeRadius, util.WithColor(n.Color))
		util.DrawText(img, n.Text, vNode, util.MplusNormalFaceSm)

		for _, e := range n.Edges {
			lineWidth := max(e.Weight*2.5, 1)
			arrowLen := max(lineWidth*2, 5)

			vAdjust := e.vUnit.Mul(NodeRadius)
			v1 := vNode.Sub(vAdjust)
			v2 := e.To.pos.EProd(r.axisMultiplier)
			v3 := v2.Add(vAdjust.Normalize().Mul(vAdjust.Len() + arrowLen))

			util.DrawLine(img, v1, v3, util.WithWidth(lineWidth))

			util.DrawFilledArrow(img, v3, v2.Add(vAdjust))
		}
	}
}
