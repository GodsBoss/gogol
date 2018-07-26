package main

import (
	"github.com/GodsBoss/conway"

	"fmt"
	"time"
)

func main() {
	topology := conway.NewRectangularTorus(60, 20)
	game := conway.NewGame(topology, conway.ConwayRule())
	conway.MiniGlider(topology, game, 0, 0)

	for {
		fmt.Println(topology.Format(game.Fields(), "0", "."))
		time.Sleep(250 * time.Millisecond)
		diff := game.Next()
		if diff.Empty() {
			break
		}
	}
}
