package main

import "strings"

// Topology defines how the fields are connected with each other.
type Topology interface {
	// Neighbours returns all fields the given field is connected to.
	Neighbours(id FieldID) FieldIDs

	// All returns all field IDs this topology contains.
	All() FieldIDs
}

type RectangularTorus struct {
	width  int
	height int
}

func NewRectangularTorus(width int, height int) *RectangularTorus {
	return &RectangularTorus{
		width:  width,
		height: height,
	}
}

func (torus RectangularTorus) Width() int {
	return torus.width
}

func (torus RectangularTorus) Height() int {
	return torus.height
}

func (torus *RectangularTorus) ID(column int, row int) FieldID {
	return FieldID(row*torus.width + column)
}

func (torus *RectangularTorus) columnRow(id FieldID) (column, row int) {
	column = int(id) % torus.width
	row = (int(id) - column) / torus.width
	return column, row
}

func (torus *RectangularTorus) All() FieldIDs {
	ids := make(FieldIDs, 0, torus.width*torus.height)
	for column := 0; column < torus.width; column++ {
		for row := 0; row < torus.height; row++ {
			ids = append(ids, torus.ID(column, row))
		}
	}
	return ids
}

func (torus *RectangularTorus) Neighbours(id FieldID) FieldIDs {
	ids := FieldIDs{}
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

func (torus *RectangularTorus) Format(fields Fields, alive, dead string) string {
	format := map[CellValue]string{
		Dead:  dead,
		Alive: alive,
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
