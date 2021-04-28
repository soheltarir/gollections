package maps

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/trees/heaps"
)

// Counter is map for counting hashable items. Sometimes called a bag or multiset.
// Elements are stored as map keys and their counters are stored as map values
type Counter struct {
	// Map to store counters.
	// The key is an interface{} which corresponds to Key() method implemented by containers.Container interfaces
	// The value is the counter
	counterMap map[interface{}]int
	// Map to store containers.Container objects w.r.t. to the Key() for faster lookups
	objMap   map[interface{}]containers.Container
	size     int
	datatype containers.Container
}

// Add increments the counter for the element provided
func (c *Counter) Add(element interface{}) {
	x := c.datatype.Validate(element)
	_, found := c.counterMap[x.Key()]
	// Update the count
	if !found {
		c.counterMap[x.Key()] = 1
		c.size++
	} else {
		c.counterMap[x.Key()]++
	}
	// Add the containers.Container to object map
	c.objMap[x.Key()] = x
}

// Subtract decrements the counter for the element provided. Counts can be reduced below zero.
func (c *Counter) Subtract(element interface{}) {
	x := c.datatype.Validate(element)
	_, found := c.counterMap[x.Key()]
	if !found {
		c.counterMap[x.Key()] = -1
		c.size++
	} else {
		c.counterMap[x.Key()]--
	}
}

// Get returns the current counter for the object provided
func (c Counter) Get(obj interface{}) int {
	element := c.datatype.Validate(obj)
	return c.counterMap[element.Key()]
}

// MostCommon lists the n most common elements and their counts from the most common to the least.
// Returns a slice of struct containing the Container and it's count
// Time Complexity: O(n)
// Space Complexity: O(n)
func (c Counter) MostCommon(n int) map[containers.Container]int {
	reverseCounterMap := make(map[int][]interface{})
	counts := make([]interface{}, 0)
	for key, value := range c.counterMap {
		_, found := reverseCounterMap[value]
		if !found {
			reverseCounterMap[value] = []interface{}{key}
			counts = append(counts, value)
		} else {
			reverseCounterMap[value] = append(reverseCounterMap[value], key)
		}
	}
	// Create a Heap
	heap := heaps.NewMaxInt(counts...)
	nLargest := heap.NLargest(n)
	result := make(map[containers.Container]int)
	for _, count := range nLargest {
		for _, element := range reverseCounterMap[containers.ToInt(count)] {
			result[c.objMap[element]] = containers.ToInt(count)
		}
	}
	return result
}

// NewCounter instantiates a new counter object with the datatype provided
func NewCounter(datatype containers.Container) *Counter {
	return &Counter{
		datatype:   datatype,
		counterMap: make(map[interface{}]int),
		objMap:     make(map[interface{}]containers.Container),
		size:       0,
	}
}

// NewIntCounter instantiates a counter object with keys as integer variables
func NewIntCounter() *Counter {
	return NewCounter(containers.IntContainer(0))
}

// NewStringCounter instantiates a counter object with keys as string variables
func NewStringCounter() *Counter {
	return NewCounter(containers.StringContainer(""))
}
