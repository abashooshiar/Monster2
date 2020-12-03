package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	dimensionsArray := [20][3]int{
		{2, 4, 3}, {1, 5, 6}, {1, 7, 8}, {1, 9, 10}, {2, 10, 11},
		{2, 7, 12}, {3, 6, 13}, {3, 9, 14}, {4, 8, 15}, {4, 5, 16},
		{5, 12, 17}, {6, 11, 18}, {7, 14, 18}, {8, 13, 19}, {9, 16, 19},
		{10, 15, 17}, {11, 16, 20}, {12, 13, 20}, {14, 15, 20}, {17, 18, 19}}

	rand.Seed(time.Now().UnixNano())
	zufall := randomint(1, 20)
	Zufallszahl := randomint(1, 20)
	fmt.Println("Zufallszahlenstart(Nur Erinnerungstest):", Zufallszahl)
	angrenzend(zufall, dimensionsArray, Zufallszahl)

}
func angrenzend(zufall int, dimensionsArray [20][3]int, Zufallszahl int) {
	fmt.Println("Du befindest dich in Raum : ", zufall)
	fmt.Println("Angrenzende Räume sind : ", dimensionsArray[zufall-1])

	for _, inhalts := range dimensionsArray[zufall-1] {
		if inhalts == Zufallszahl {
			fmt.Println("Es stinkt bestialisch.")
		}
	}

	fmt.Print("Wohin möchtest Du gehen? : ")
	var wohingehen int
	if _, err := fmt.Scan(&wohingehen); err != nil {
		fmt.Print("Bitte geben Sie eine Nummer__ ")
	}

	if wohingehen == Zufallszahl {
		fmt.Println("Du hast das Monster gefunden!")
		fmt.Println("drücken Sie 'j', um das Spiel zu beenden.Andernfalls geben Sie eine Zahl ein.")

		var exit int
		if _, err := fmt.Scan(&exit); err != nil {
			os.Exit(0)
		}
	}
	finder(zufall, wohingehen, dimensionsArray, Zufallszahl)
}

func randomint(min, max int) int {
	return min + rand.Intn(max-min)
}
func finder(zufall, wohingehen int, dimensional [20][3]int, Zufallszahl int) {
	indexOfDimensional := dimensional[zufall-1]

	sw, _ := findin(indexOfDimensional[:], wohingehen, Zufallszahl)
	if sw == true {
		fmt.Println("----------------------------")
		fmt.Println("Zufallszahl(Erinnerungstest und anknüpfen):", Zufallszahl)
		zufall = wohingehen
		angrenzend(zufall, dimensional, Zufallszahl)
	} else {
		fmt.Println("Du kannst von hier nicht in Raum ", wohingehen)
		angrenzend(zufall, dimensional, Zufallszahl)
	}
}
func findin(arrayinhalt []int, wohingehen int, Zufallszahl int) (bool, int) {
	switchKey := false
	for _, inhalt := range arrayinhalt {
		if inhalt == wohingehen {
			switchKey = true
			//return switchKey, Zufallszahl
		}
	}
	return switchKey, 0
}
