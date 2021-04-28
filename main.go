package main

import (
	"fmt"
	"github.com/soheltarir/gollections/maps"
)

func main() {
	counter := maps.NewStringCounter()
	counter.Add("sohel")
	counter.Add("tarir")
	counter.Add("sohel")
	fmt.Println(counter.MostCommon(2))
}
