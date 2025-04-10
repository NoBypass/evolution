package graph

import (
	"evolution/internal/util"
	"image/color"
	"math/rand/v2"
)

const (
	graphWidth, graphHeight = 1000, 1000 // bounds for the simulation

	centerX, centerY = graphWidth / 2, graphHeight / 2

	kSpring    = 0.1  // stiffness of each spring (edge)
	kRepulsion = 1000 // repulsion strength of the nodes
	kCentering = 1e-4 // force to keep nodes near the center

	timeStep = 0.02 // controls animation speed/stability

	desiredDistance = 150 // the desired length of each spring (edge)
)

var (
	timeStepVec   = util.Vec2(timeStep, timeStep)
	centerVec     = util.Vec2(centerX, centerY)
	kCenteringVec = util.Vec2(kCentering, kCentering)
)

type Node struct {
	pos      util.Vector2
	velocity util.Vector2
	force    util.Vector2
	Text     string
	Edges    []*Edge
	Color    color.Color
}

type Edge struct {
	To     *Node
	Weight float32
	vUnit  util.Vector2
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) tick() {
	// Reset forces
	for _, n := range g.Nodes {
		n.force = util.Vec2(0, 0)
	}

	// Repulsion forces
	for i, a := range g.Nodes {
		for j := i + 1; j < len(g.Nodes); j++ {
			b := g.Nodes[j]
			distVec := a.pos.Sub(b.pos)
			dsq := distVec.Dot(distVec)

			if dsq < 1e-8 {
				distVec = util.Vec2(rand.Float32()*2-1, rand.Float32()*2-1)
				dsq = distVec.Dot(distVec)
			}

			forceVec := distVec.Mul(kRepulsion / dsq).Div(distVec.Len())

			a.force = a.force.Add(forceVec)
			b.force = b.force.Sub(forceVec)
		}
	}

	// Spring (edge attraction) forces
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			to := e.To
			distVec := n.pos.Sub(to.pos)
			dist := distVec.Len()
			if dist == 0 {
				continue // avoid division by zero
			}

			e.vUnit = distVec.Normalize()

			forceVec := distVec.Mul(-kSpring * (dist - desiredDistance)).Div(dist)

			n.force = n.force.Add(forceVec)
			to.force = to.force.Sub(forceVec)
		}
	}

	// Centering forces
	for _, n := range g.Nodes {
		n.force = n.force.Add(kCenteringVec.EProd(centerVec.Sub(n.pos)))
	}

	// Apply forces to velocity and update positions
	for _, n := range g.Nodes {
		n.velocity = n.velocity.Add(timeStepVec.EProd(n.force))

		n.pos = n.pos.Add(timeStepVec.EProd(n.velocity))
		n.pos = util.Vec2(min(max(n.pos.X, 0), graphWidth), min(max(n.pos.Y, 0), graphHeight))

		n.velocity = n.velocity.Mul(0.85)
	}
}
