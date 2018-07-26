package gogol

// Topology defines how the fields are connected with each other.
type Topology interface {
	// Neighbours returns all fields the given field is connected to.
	Neighbours(id FieldID) FieldIDs

	// All returns all field IDs this topology contains.
	All() FieldIDs
}
