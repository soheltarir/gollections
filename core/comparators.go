package core

// Comparator function compares two asserted types.
// This functions implemented using this signature will panic, if concerned type assertions fail.
// See IntCompare function for reference
type Comparator func(x, y interface{}) ComparatorResult

// IntCompare provides comparison between int datatypes
func IntCompare(x, y interface{}) ComparatorResult {
	first := x.(int)
	second := y.(int)
	if first < second {
		return Lesser
	} else if first > second {
		return Greater
	} else {
		return Equal
	}
}

// StringCompare provides comparison between string data types
func StringCompare(x, y interface{}) ComparatorResult {
	first := x.(string)
	second := y.(string)
	if first < second {
		return Lesser
	} else if first > second {
		return Greater
	} else {
		return Equal
	}
}
