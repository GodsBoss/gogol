package gogol

// History stores a game's history.
type History struct {
	// N is the amount of game states stored.
	N int

	// States contains game states.
	States []Fields
}

// Add adds a game state to the history.
func (history *History) Add(fields Fields) {
	if len(history.States) >= history.N {
		history.States = history.States[1:]
	}
	history.States = append(history.States, fields)
}

// CellValue returns the cell values for a field identified by id from oldest to newest.
func (history *History) CellValue(id FieldID) CellValueHistory {
	result := make(CellValueHistory, len(history.States))
	for i := range history.States {
		result[i] = history.States[i].At(id)
	}
	return result
}

// CellValueHistory is a list cell values, from oldest to newest.
type CellValueHistory []CellValue

func (history CellValueHistory) countCellValue(value CellValue) int {
	count := 0
	for i := range history {
		if history[i] == value {
			count++
		}
	}
	return count
}

// AliveCount returns how often a cell was alive.
func (history CellValueHistory) AliveCount() int {
	return history.countCellValue(Alive)
}

// DeadCount returns how often a cell was dead.
func (history CellValueHistory) DeadCount() int {
	return history.countCellValue(Dead)
}

// Count returns both how often a cell was alive and dead.
func (history CellValueHistory) Count() (alive, dead int) {
	return history.AliveCount(), history.DeadCount()
}
