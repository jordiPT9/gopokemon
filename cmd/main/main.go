package main

import (
	"fmt"

	"github.com/jordipt9/gopokemon/database"
)

func main() {
	database.Init()
	list := []string{
		"Slowking",
		"Misdreavus",
		"Unown",
		"Wobbuffet",
		"Girafarig",
		"Pineco",
		"Forretress",
		"Dunsparce",
		"Gligar",
		"Steelix",
		"Snubbull",
		"Granbull",
		"Qwilfish",
		"Scizor",
		"Shuckle",
		"Heracross",
		"Sneasel",
		"Teddiursa",
		"Ursaring",
		"Slugma",
		"Magcargo",
		"Swinub",
		"Piloswine",
		"Corsola",
		"Remoraid",
		"Octillery",
		"Delibird",
		"Mantine",
		"Skarmory",
		"Houndour",
		"Houndoom",
		"Kingdra",
		"Phanpy",
		"Donphan",
		"Porygon2",
		"Stantler",
		"Smeargle",
		"Tyrogue",
		"Hitmontop",
		"Smoochum",
	}
	resultList, err := database.TopRawMaxAtk(list)
	if err != nil {
		panic(err)
	}
	for _, pkmn := range resultList {
		fmt.Println(pkmn.Name)
	}
	fmt.Println("RAW ATK: ", resultList)

	fmt.Println("")
	resultList, err = database.TopBalancedMaxAtk(list)
	if err != nil {
		panic(err)
	}
	for _, pkmn := range resultList {
		fmt.Println(pkmn.Name)
	}
	fmt.Println("BALANCED ATK: ", resultList)
}
