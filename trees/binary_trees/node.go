package binary_trees

import "github.com/soheltarir/gollections/containers"

type Node struct {
	Value containers.Container
	Left  *Node
	Right *Node
}

func (n Node) Key() interface{} {
	return n.Value.Key()
}

func (n Node) Less(x containers.Container) bool {
	y := x.(Node)
	return n.Value.Less(y.Value)

}

func (Node) Validate(x interface{}) containers.Container {
	converted, ok := x.(Node)
	if !ok {
		return x.(*Node)
	}
	return converted
}
