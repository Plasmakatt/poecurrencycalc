package main

import (
	"fmt"

	"github.com/plasmakatt/poecurrencycalc/poeninja"
)

func main() {
	var poeNinja = poeninja.PoeNinja{}
	poeNinja.LoadData()
	for _, items := range poeNinja.Items {
		fmt.Println("Item: " + items.Name)
		fmt.Println("Worth:" + fmt.Sprintf("%f", items.ChaosValue) + "C")
	}
}
