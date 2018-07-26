package gogol

type Rule interface {
	NextValue(current CellValue, alive int, dead int) CellValue
}

type RuleFunc func(current CellValue, alive int, dead int) CellValue

func (f RuleFunc) NextValue(current CellValue, alive int, dead int) CellValue {
	return f(current, alive, dead)
}

// AliveMappingRule maps the count of cells which are alive to a new cell value.
// Missing values mean the current value is kept.
type AliveMappingRule map[CellValue]map[int]CellValue

func (m AliveMappingRule) NextValue(current CellValue, alive int, _ int) CellValue {
	if next, ok := m[current][alive]; ok {
		return next
	}
	return current
}
