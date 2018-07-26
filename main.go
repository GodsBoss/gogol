package conway

import (
	"math/rand"
	"time"
)

func viveRandomFields(topology Topology, game *Game, count int) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	all := topology.All()
	for i := 0; i < count; i++ {
		game.Override(all[random.Intn(len(all))], Alive)
	}
}

func MiniGlider(topology *RectangularTorus, game *Game, c, r int) {
	coords := []struct {
		X int
		Y int
	}{
		{X: c, Y: r + 1},
		{X: c, Y: r + 3},
		{X: c + 1, Y: r},
		{X: c + 2, Y: r},
		{X: c + 3, Y: r},
		{X: c + 4, Y: r},
		{X: c + 4, Y: r + 1},
		{X: c + 4, Y: r + 2},
		{X: c + 3, Y: r + 3},
	}
	for i := range coords {
		game.Override(topology.ID(coords[i].X, coords[i].Y), Alive)
	}
}
