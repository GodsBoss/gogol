package gogol

// Game is a Game of Life instance. It contains a topology, a rule and the
// current fields.
// Must be instantiated with NewGame, the zero value is useless.
type Game struct {
	topology Topology
	rule     Rule

	currentFields Fields
	nextFields    Fields
}

// NewGame creates a new game of life with the given topology and rule. Sets all
// cell values to Dead.
func NewGame(topology Topology, rule Rule) *Game {
	game := &Game{
		topology:      topology,
		rule:          rule,
		currentFields: Fields{},
		nextFields:    Fields{},
	}

	all := topology.All()
	for i := range all {
		game.currentFields[all[i]] = Dead
		game.nextFields[all[i]] = Dead
	}

	return game
}

// Next advances the game to the next state. Returns the difference to the previous state.
func (game *Game) Next() Fields {
	diff := Fields{}
	for fieldID := range game.currentFields {
		alive, dead := game.topology.Neighbours(fieldID).Count(game.currentFields)
		currentValue := game.currentFields.At(fieldID)
		nextValue := game.rule.NextValue(game.currentFields.At(fieldID), alive, dead)
		if currentValue != nextValue {
			diff[fieldID] = nextValue
		}
		nextValue.apply(game.nextFields, fieldID)
	}
	game.currentFields, game.nextFields = game.nextFields, game.currentFields
	return diff
}

// Fields returns the current states of all cells. Must not be changed.
func (game *Game) Fields() Fields {
	return game.currentFields
}

// Override sets the cell at the given index to a new value.
func (game *Game) Override(id FieldID, value CellValue) {
	game.currentFields.Set(id, value)
}
