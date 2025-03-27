package game

import "evolution/internal/environment"

func (g *Game) getOrganismAtCoordinates(x, y float32) *environment.Organism {
	cellSize := float32(g.WindowSize) / float32(g.currentEnv.Size)

	for _, org := range g.currentEnv.Organisms {
		orgX := float32(org.X) * cellSize
		orgY := float32(org.Y)*cellSize + HeaderOffset

		if x >= orgX && x <= orgX+cellSize && y >= orgY && y <= orgY+cellSize {
			return org
		}
	}

	return nil
}
