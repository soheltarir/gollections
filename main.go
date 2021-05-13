package main

import (
	"fmt"
	"github.com/soheltarir/gollections/lists"
)

func main() {
	ll := lists.NewInt()
	//var data = [][]interface{}{
	//	{1, 2, 3, 4, 5, 6},
	//	{1, 2, 3, 4, 5, 6},
	//	{1, 2, 3, 4, 5, 6},
	//	{1, 2, 3, 4, 5, 6},
	//	{1, 2, 3, 4, 5, 8},
	//	{1, 2, 3, 4, 5, 7},
	//}
	//var wg sync.WaitGroup
	//for index, datum := range data {
	//	wg.Add(1)
	//	go func(datum []interface{}, index int) {
	//		defer wg.Done()
	//		fmt.Printf("Executing goroutine %d, Begin: %v, Data: %v\n", index, ll.Begin(), datum)
	//		for _, ele := range datum {
	//			ll.PushBack(ele)
	//		}
	//		// ll.Insert(ll.Begin(), datum...)
	//	}(datum, index)
	//}
	//wg.Wait()
	ll.Insert(ll.Begin(), 1, 2, 3, 4, 5, 6)
	fmt.Println("Total Size: ", ll.Size())
	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		fmt.Printf("%d -> ", it.Value())
	}
}
