package containers_test

import (
	"fmt"
	"github.com/soheltarir/gollections/containers"
)

func ExampleToInt() {
	a := containers.IntContainer(1)
	// Convert container to basic int type
	fmt.Printf("Value: %d Type: %T", containers.ToInt(a), containers.ToInt(a))

	// The below code will panic
	containers.ToInt("a")

	// Output:
	// Value: 1 Type: int
	// panic: invalid type provided; expected: containers.IntContainer, received: string [recovered]
	//        panic: invalid type provided; expected: containers.IntContainer, received: string
}

func ExampleToString() {
	a := containers.StringContainer("a")
	// Convert container to basic int type
	fmt.Printf("Value: %s Type: %T", containers.ToString(a), containers.ToString(a))

	// The below code will panic
	containers.ToString(1)

	// Output:
	// Value: a Type: string
	// panic: invalid type provided; expected: containers.StringContainer, received: int [recovered]
	//        panic: invalid type provided; expected: containers.StringContainer, received: int
}
