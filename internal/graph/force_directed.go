package graph

import (
	"github.com/chewxy/math32"
	"image/color"
	"math/rand/v2"
)

const (
	graphWidth, graphHeight = 1000, 1000 // bounds for the simulation
	centerX, centerY        = graphWidth / 2, graphHeight / 2

	kSpring    = 0.1  // stiffness of each spring (edge)
	kRepulsion = 1000 // repulsion strength of the nodes
	kCentering = 1e-4 // force to keep nodes near the center

	timeStep = 0.02 // controls animation speed/stability

	desiredDistance = 150 // the desired length of each spring (edge)
)

type Node struct {
	X, Y   float32 // position
	VX, VY float32 // velocity
	FX, FY float32 // force
	Text   string
	Edges  []*Edge
	Color  color.Color
}

type Edge struct {
	To     *Node
	Weight float32
	ux, uy float32 // partial unit vector from parent to child node
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) tick() {
	// Reset forces
	for _, n := range g.Nodes {
		n.FX, n.FY = 0, 0
	}

	// Repulsion forces
	for i, a := range g.Nodes {
		for j := i + 1; j < len(g.Nodes); j++ {
			b := g.Nodes[j]
			dx, dy := a.X-b.X, a.Y-b.Y
			dsq := dx*dx + dy*dy

			if dsq < 1e-8 {
				dx, dy = rand.Float32()*2-1, rand.Float32()*2-1
				dsq = dx*dx + dy*dy
			}

			dist := math32.Sqrt(dsq)
			force := kRepulsion / dsq
			fx, fy := force*dx/dist, force*dy/dist

			a.FX += fx
			a.FY += fy
			b.FX -= fx
			b.FY -= fy
		}
	}

	// Spring (edge attraction) forces
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			to := e.To
			dx, dy := n.X-to.X, n.Y-to.Y
			dist := math32.Sqrt(dx*dx + dy*dy)
			if dist == 0 {
				continue // avoid division by zero
			}

			e.ux, e.uy = dx/dist, dy/dist

			force := -kSpring * (dist - desiredDistance)
			fx, fy := force*dx/dist, force*dy/dist

			n.FX += fx
			n.FY += fy
			to.FX -= fx
			to.FY -= fy
		}
	}

	// Centering forces
	for _, n := range g.Nodes {
		dx := centerX - n.X
		dy := centerY - n.Y
		n.FX += kCentering * dx
		n.FY += kCentering * dy
	}

	// Apply forces to velocity and update positions
	for _, n := range g.Nodes {
		n.VX += timeStep * n.FX
		n.VY += timeStep * n.FY

		n.X += timeStep * n.VX
		n.Y += timeStep * n.VY

		n.VX *= 0.85
		n.VY *= 0.85

		n.X = math32.Min(math32.Max(n.X, 0), graphWidth)
		n.Y = math32.Min(math32.Max(n.Y, 0), graphHeight)
	}
}
