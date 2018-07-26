package gogol

import (
	"math"
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

// FillWithRandomFields fills a game with cells randomly set to alive or dead with
// a certain distribution.
func FillWithRandomFields(game *Game, aliveShare float64) {
	aliveShare = math.Min(1, math.Max(0, aliveShare))
	origAll := game.topology.All()
	all := make(FieldIDs, len(origAll))
	copy(all, origAll)
	random := rand.New(rand.NewSource(time.Now().Unix()))
	random.Shuffle(
		len(all),
		func(i, j int) {
			all[i], all[j] = all[j], all[i]
		},
	)
	pivot := int(aliveShare * float64(len(all)))
	for i := 0; i < pivot; i++ {
		game.Override(all[i], Alive)
	}
	for i := pivot; i < len(all); i++ {
		game.Override(all[i], Dead)
	}
}
