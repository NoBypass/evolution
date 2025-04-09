package graph

import (
	"math"
	"math/rand/v2"
)

const (
	Width, Height = 800, 800 // bounds for the simulation

	kSpring    = 0.1   // stiffness of each spring (edge)
	kRepulsion = 10000 // repulsion strength of the nodes

	timeStep = 0.02 // controls animation speed/stability

	minDistance     = 0.1 // the minimal distance between each node (to prevent divide-by-zero)
	desiredDistance = 100 // the desired length of each spring (edge)
)

type Node struct {
	X, Y   float64 // position
	VX, VY float64 // velocity
	FX, FY float64 // force
	Edges  []*Edge
}

type Edge struct {
	To     *Node
	Weight float64
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) tick() {
	for _, n := range g.Nodes {
		n.FX, n.FY = 0, 0
	}

	for i, a := range g.Nodes {
		for j := i + 1; j < len(g.Nodes); j++ {
			b := g.Nodes[j]
			dx := a.X - b.X
			dy := a.Y - b.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < 1e-4 {
				dx = rand.Float64()*2 - 1
				dy = rand.Float64()*2 - 1
				dist = math.Sqrt(dx*dx + dy*dy)
			}
			force := kRepulsion / (dist * dist)

			fx := force * dx / dist
			fy := force * dy / dist

			a.FX += fx
			a.FY += fy
			b.FX -= fx
			b.FY -= fy
		}
	}

	for _, node := range g.Nodes {
		for _, edge := range node.Edges {
			to := edge.To
			dx := node.X - to.X
			dy := node.Y - to.Y
			dist := math.Sqrt(dx*dx+dy*dy) + minDistance
			force := -kSpring * (dist - desiredDistance)

			fx := force * dx / dist
			fy := force * dy / dist

			node.FX += fx
			node.FY += fy
			to.FX -= fx
			to.FY -= fy
		}
	}

	for _, n := range g.Nodes {
		n.VX += timeStep * n.FX
		n.VY += timeStep * n.FY

		n.X += timeStep * n.VX
		n.Y += timeStep * n.VY

		n.VX *= 0.85
		n.VY *= 0.85
		n.X = math.Min(math.Max(n.X, 0), Width)
		n.Y = math.Min(math.Max(n.Y, 0), Height)
	}
}
