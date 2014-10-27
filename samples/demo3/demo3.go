package main

import (
	"fmt"

	"github.com/johannesboyne/mor-gopresentation/samples/demo3/mor"
)

func main() {
	rob := mor.NewGopher("Rob", 26)
	peter := mor.NewGopher("Peter", 24)
	mike := mor.NewGopher("Mike", 31)
	johannes := mor.MoRer{"Johannes", 23, "johannes@boyne.de"}
	rob.Talk()

	fmt.Printf("total age: %v", totalAge(&rob, &peter, &mike, &johannes))
}

func totalAge(heros ...mor.Hero) int {
	tage := 0
	for _, h := range heros {
		tage += h.GetAge()
	}
	return tage
}
