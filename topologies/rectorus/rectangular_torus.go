package rectorus

import (
	"github.com/GodsBoss/gogol"

	"strings"
)

type Topology struct {
	width  int
	height int
}

func NewRectangularTorus(width int, height int) *Topology {
	return &Topology{
		width:  width,
		height: height,
	}
}

func (torus Topology) Width() int {
	return torus.width
}

func (torus Topology) Height() int {
	return torus.height
}

func (torus *Topology) ID(column int, row int) gogol.FieldID {
	return gogol.FieldID(row*torus.width + column)
}

func (torus *Topology) columnRow(id gogol.FieldID) (column, row int) {
	column = int(id) % torus.width
	row = (int(id) - column) / torus.width
	return column, row
}

func (torus *Topology) All() gogol.FieldIDs {
	ids := make(gogol.FieldIDs, 0, torus.width*torus.height)
	for column := 0; column < torus.width; column++ {
		for row := 0; row < torus.height; row++ {
			ids = append(ids, torus.ID(column, row))
		}
	}
	return ids
}

func (torus *Topology) Neighbours(id gogol.FieldID) gogol.FieldIDs {
	ids := gogol.FieldIDs{}
	centerColumn, centerRow := torus.columnRow(id)
	for columnOffset := -1; columnOffset <= 1; columnOffset++ {
		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
			if columnOffset == 0 && rowOffset == 0 {
				continue
			}
			column := (centerColumn + columnOffset + torus.width) % torus.width
			row := (centerRow + rowOffset + torus.height) % torus.height
			ids = append(ids, torus.ID(column, row))
		}
	}
	return ids
}

func (torus *Topology) Format(fields gogol.Fields, alive, dead string) string {
	format := map[gogol.CellValue]string{
		gogol.Dead:  dead,
		gogol.Alive: alive,
	}
	parts := make([]string, 0, torus.height*(torus.width+1))
	for row := 0; row < torus.height; row++ {
		for column := 0; column < torus.width; column++ {
			parts = append(parts, format[fields.At(torus.ID(column, row))])
		}
		parts = append(parts, "\n")
	}
	return strings.Join(parts, "")
}

func MiniGlider(topology *Topology, game *gogol.Game, c, r int) {
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
		game.Override(topology.ID(coords[i].X, coords[i].Y), gogol.Alive)
	}
}
