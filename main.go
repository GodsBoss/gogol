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
