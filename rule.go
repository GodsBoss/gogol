package gogol

// Rule rule abstracts how a cell changes its state in the next generation.
type Rule interface {
	// NextValue calculates the next cell value from the current value and the
	// amounts of alive and dead cells around it.
	NextValue(current CellValue, alive int, dead int) CellValue
}

// RuleFunc implements Rule via a function with the same signature as Rule.NextValue.
type RuleFunc func(current CellValue, alive int, dead int) CellValue

// NextValue returns the result of passing all parameters to f.
func (f RuleFunc) NextValue(current CellValue, alive int, dead int) CellValue {
	return f(current, alive, dead)
}

// AliveMappingRule maps the count of cells which are alive to a new cell value.
// Missing values mean the current value is kept.
type AliveMappingRule map[CellValue]map[int]CellValue

// NextValue searches for a matching entry in the map and returns a corresponding value.
// If no value is found, the current value is returned.
func (m AliveMappingRule) NextValue(current CellValue, alive int, _ int) CellValue {
	if next, ok := m[current][alive]; ok {
		return next
	}
	return current
}
