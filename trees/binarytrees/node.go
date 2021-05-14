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

package binarytrees

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
