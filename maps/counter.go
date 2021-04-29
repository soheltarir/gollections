package maps

import (
	"github.com/soheltarir/gollections/containers"
	"github.com/soheltarir/gollections/trees/heaps"
	"sync"
)

// Counter is map for counting hashable items. Sometimes called a bag or multiset.
// Elements are stored as map keys and their counters are stored as map values
type Counter struct {
	// Map to store counters.
	// The key is an interface{} which corresponds to Key() method implemented by containers.Container interfaces
	// The value is the counter
	countMap sync.Map
	// Map to store containers.Container objects w.r.t. to the Key() for faster lookups
	objectMap sync.Map
	size      int
	datatype  containers.Container
}

func (c *Counter) Size() int {
	return c.size
}

func (c *Counter) _getFromCountMap(key interface{}) (int, bool) {
	value, found := c.countMap.Load(key)
	if !found {
		return 0, found
	}
	return value.(int), found
}

func (c *Counter) _storeInCountMap(key interface{}, value int) {
	c.countMap.Store(key, value)
}

func (c *Counter) _storeInObjectMap(key interface{}, value containers.Container) {
	c.objectMap.Store(key, value)
}

func (c *Counter) _getFromObjectMap(key interface{}) (containers.Container, bool) {
	value, found := c.objectMap.Load(key)
	return value.(containers.Container), found
}

// Add increments the counter for the element provided
func (c *Counter) Add(element interface{}) {
	x := c.datatype.Validate(element)
	currCounter, found := c._getFromCountMap(x.Key())
	// Update the count
	if !found {
		c._storeInCountMap(x.Key(), 1)
		c.size++
	} else {
		currCounter++
		c._storeInCountMap(x.Key(), currCounter)
	}
	// Add the containers.Container to object map
	c._storeInObjectMap(x.Key(), x)
}

// AddMany updates the counts for the arguments provided
func (c *Counter) AddMany(elements ...interface{}) {
	for _, element := range elements {
		c.Add(element)
	}
}

// Subtract decrements the counter for the element provided. Counts can be reduced below zero.
func (c *Counter) Subtract(element interface{}) {
	x := c.datatype.Validate(element)
	currCount, found := c._getFromCountMap(x.Key())
	if !found {
		c._storeInCountMap(x.Key(), -1)
		c.size++
	} else {
		currCount--
		c._storeInCountMap(x.Key(), currCount)
	}
}

// Delete removes an item from the counter map completely. If counter is nil or there is no such element, delete
// is a no-op.
func (c *Counter) Delete(element interface{}) {
	x := c.datatype.Validate(element)
	c.countMap.Delete(x.Key())
	c.objectMap.Delete(x.Key())
	if c.size > 0 {
		c.size--
	}
}

// Get returns the current counter for the object provided. Returns zero if the key is not found in the counter
func (c *Counter) Get(obj interface{}) int {
	element := c.datatype.Validate(obj)
	count, found := c._getFromCountMap(element.Key())
	if !found {
		return 0
	}
	return count
}

// Range calls callback sequentially for each key & value (the counter) in the Counter object.
// This function internally uses sync.Map's Range method, and hence can show inconsistencies during concurrency.
func (c *Counter) Range(callback func(key interface{}, value int)) {
	c.countMap.Range(func(k, v interface{}) bool {
		callback(k, v.(int))
		return true
	})
}

// MostCommon lists the n most common elements and their counts from the most common to the least.
// Returns a slice of struct containing the Container and it's count
// Time Complexity: O(n)
// Space Complexity: O(n)
func (c *Counter) MostCommon(n int) map[containers.Container]int {
	reverseCounterMap := make(map[int][]interface{})
	counts := make([]interface{}, 0)
	c.countMap.Range(func(key, value interface{}) bool {
		valInt := value.(int)
		_, found := reverseCounterMap[valInt]
		if !found {
			reverseCounterMap[value.(int)] = []interface{}{key}
			counts = append(counts, valInt)
		} else {
			reverseCounterMap[valInt] = append(reverseCounterMap[valInt], key)
		}
		return true
	})
	// Create a Heap
	heap := heaps.NewMaxInt(counts...)
	nLargest := heap.NLargest(n)
	result := make(map[containers.Container]int)
	for _, count := range nLargest {
		for _, element := range reverseCounterMap[containers.ToInt(count)] {
			obj, _ := c._getFromObjectMap(element)
			result[obj] = containers.ToInt(count)
		}
	}
	return result
}

// NewCounter instantiates a new counter object with the datatype provided
func NewCounter(datatype containers.Container, elements ...interface{}) *Counter {
	counter := &Counter{
		datatype: datatype,
		size:     0,
	}
	if len(elements) > 0 {
		counter.AddMany(elements...)
	}
	return counter
}

// NewIntCounter instantiates a counter object with keys as integer variables
func NewIntCounter(elements ...interface{}) *Counter {
	return NewCounter(containers.IntContainer(0), elements...)
}

// NewStringCounter instantiates a counter object with keys as string variables
func NewStringCounter(elements ...interface{}) *Counter {
	return NewCounter(containers.StringContainer(""), elements...)
}
