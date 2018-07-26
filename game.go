package gogol

type Game struct {
	topology Topology
	rule     Rule

	currentFields Fields
	nextFields    Fields
}

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

func (game *Game) Fields() Fields {
	return game.currentFields
}

func (game *Game) Override(id FieldID, value CellValue) {
	game.currentFields.Set(id, value)
}
