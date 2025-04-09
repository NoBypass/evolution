package graph

import (
	"evolution/internal/environment"
	"evolution/internal/environment/neural"
	"evolution/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	// axis multiplier
	mx, my float32
}

func NewRenderer(img *ebiten.Image) *Renderer {
	return &Renderer{
		img:   img,
		graph: &Graph{},

		OrgCh: make(chan *environment.Organism),
		mx:    float32(img.Bounds().Max.X) / graphWidth,
		my:    float32(img.Bounds().Max.Y) / graphHeight,
	}
}

func (r *Renderer) Run() {
	org := <-r.OrgCh
	for {
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
	var clr color.Color = color.Black
	switch neuron.Type {
	case neural.Sensory:
		clr = color.RGBA{94, 233, 181, 255}
	case neural.Internal:
		clr = color.RGBA{212, 212, 212, 255}
	case neural.Action:
		clr = color.RGBA{255, 184, 106, 255}
	}
	node := &Node{
		X:     graphWidth / 2,
		Y:     graphHeight / 2,
		Text:  neuron.ID.String(),
		Color: clr,
	}
	for _, s := range neuron.Incoming {
		var to *Node
		if n, ok := m[neuron]; ok {
			to = n
		} else {
			to = traverse(s.From, m)
		}

		node.Edges = append(node.Edges, &Edge{
			Weight: s.Weight,
			To:     to,
		})
	}

	m[neuron] = node
	return node
}

func (r *Renderer) Draw(img *ebiten.Image) {
	img.Clear()
	for _, n := range r.graph.Nodes {
		x, y := n.X*r.mx, n.Y*r.my
		vector.DrawFilledCircle(
			img,
			x,
			y,
			NodeRadius,
			n.Color,
			true)
		drawText(img, n.Text, x, y)
		for _, e := range n.Edges {
			vector.StrokeLine(
				img,
				x-NodeRadius*e.ux,
				y-NodeRadius*e.uy,
				e.To.X*r.my+NodeRadius*e.ux,
				e.To.Y*r.my+NodeRadius*e.uy,
				max(e.Weight*2.5, 1),
				color.Black,
				true)
		}
	}
}

func drawText(img *ebiten.Image, str string, x, y float32) {
	op := &text.DrawOptions{}
	op.ColorScale.Scale(0, 0, 0, 1)
	op.GeoM.Translate(float64(x), float64(y))
	op.LayoutOptions.PrimaryAlign = text.AlignCenter
	op.LayoutOptions.SecondaryAlign = text.AlignCenter
	text.Draw(img, str, utils.MplusNormalFaceSm, op)
}
