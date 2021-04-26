package core

// TypeChecker function does type assertions
type TypeChecker func(x interface{}) bool

// IntChecker checks whether the interface provided is an int data.
func IntChecker(x interface{}) bool {
	_, ok := x.(int)
	return ok
}

// StringChecker checks whether the interface provided is a string data.
func StringChecker(x interface{}) bool {
	_, ok := x.(string)
	return ok
}
