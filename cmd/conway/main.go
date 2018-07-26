package main

import (
	"github.com/GodsBoss/conway"
	"github.com/GodsBoss/conway/topologies/rectorus"

	"fmt"
	"time"
)

func main() {
	topology := rectorus.NewRectangularTorus(60, 20)
	game := conway.NewGame(topology, conway.ConwayRule())
	rectorus.MiniGlider(topology, game, 0, 0)

	for {
		fmt.Println(topology.Format(game.Fields(), "0", "."))
		time.Sleep(250 * time.Millisecond)
		diff := game.Next()
		if diff.Empty() {
			break
		}
	}
}
