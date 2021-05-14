/**
MIT License

Copyright (c) 2021 Sohel Tarir

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// Package containers exposes interfaces to build containers, the basic building block for gollections data-structures.
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
