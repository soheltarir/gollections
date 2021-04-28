package containers

// ToInt converts a Container object to basic integer type. Panics if unexpected type is received.
func ToInt(value interface{}) int {
	converted, ok := value.(IntContainer)
	if !ok {
		panic(buildErrorMsg(IntContainer(0), value))
	}
	return int(converted)
}

// ToIntSlice converts slice of Container objects to slice of basic integer slice
func ToIntSlice(values []Container) []int {
	var converted []int
	for _, value := range values {
		converted = append(converted, ToInt(value))
	}
	return converted
}

// ToString converts a Container object to basic string type. Panics if unexpected type is received.
func ToString(value interface{}) string {
	converted, ok := value.(StringContainer)
	if !ok {
		panic(buildErrorMsg(StringContainer(""), value))
	}
	return string(converted)
}

// ToStringSlice converts slice of Container objects to slice of basic string slice
func ToStringSlice(values []Container) []string {
	var converted []string
	for _, value := range values {
		converted = append(converted, ToString(value))
	}
	return converted
}
