package conway

import (
	"fmt"
)

type FieldID int

type FieldIDs []FieldID

func (ids FieldIDs) Count(fields Fields) (alive, dead int) {
	return ids.AliveCount(fields), ids.DeadCount(fields)
}

func (ids FieldIDs) AliveCount(fields Fields) int {
	return countFieldCellValues(Alive, ids, fields)
}

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

type Fields map[FieldID]CellValue

func (f Fields) Empty() bool {
	return len(f) == 0
}

func (f Fields) At(id FieldID) CellValue {
	return f[id]
}

func (f Fields) Set(id FieldID, value CellValue) {
	f[id] = value
}

func (f Fields) Kill(id FieldID) {
	f[id] = Dead
}

func (f Fields) Vive(id FieldID) {
	f[id] = Alive
}

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

const (
	Dead CellValue = iota
	Alive
)
