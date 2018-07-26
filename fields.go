package gogol

import (
	"fmt"
)

// FieldID is the index of a field.
type FieldID int

// FieldIDs represents a bunch of field indexes.
type FieldIDs []FieldID

// Count returns both the alive and dead cell counts.
func (ids FieldIDs) Count(fields Fields) (alive, dead int) {
	return ids.AliveCount(fields), ids.DeadCount(fields)
}

// AliveCount returns the number of cells currently alive.
func (ids FieldIDs) AliveCount(fields Fields) int {
	return countFieldCellValues(Alive, ids, fields)
}

// DeadCount returns the number of cells currently dead.
func (ids FieldIDs) DeadCount(fields Fields) int {
	return countFieldCellValues(Dead, ids, fields)
}

func countFieldCellValues(value CellValue, ids FieldIDs, fields Fields) int {
	result := 0
	for i := range ids {
		if value == fields.At(ids[i]) {
			result++
		}
	}
	return result
}

// Fields maps field indexes to cell values.
type Fields map[FieldID]CellValue

// Empty checks wether f contains nothing.
func (f Fields) Empty() bool {
	return len(f) == 0
}

// At returns the cell value for the field identified by the index.
func (f Fields) At(id FieldID) CellValue {
	return f[id]
}

// Set sets the cell value for a given index.
func (f Fields) Set(id FieldID, value CellValue) {
	f[id] = value
}

// Kill kills the cell for a given index.
func (f Fields) Kill(id FieldID) {
	f[id] = Dead
}

// Vive sets the cell for a given index to alive.
func (f Fields) Vive(id FieldID) {
	f[id] = Alive
}

// CellValue represents the value of a cell. Either dead or alive.
type CellValue int

func (value CellValue) apply(f Fields, id FieldID) {
	switch value {
	case Dead:
		f.Kill(id)
	case Alive:
		f.Vive(id)
	default:
		panic(fmt.Sprintf("invalid cell value %d", value))
	}
}

// Cell values.
const (
	Dead CellValue = iota
	Alive
)
