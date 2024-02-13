package main

import (
	"fmt"
)
//dont worry about that red line its just a red line
func main() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"maharashtra":  12242441423,
		"tamilnadu":    43434343434,
		"andrapradesh": 43523545345,
	}

	fmt.Println(statePopulation)
	delete(statePopulation, "maharashtra")

}
