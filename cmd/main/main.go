package main

import (
	"fmt"

	"github.com/jordipt9/gopokemon/database"
	"github.com/jordipt9/gopokemon/types"
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
	resultList, err := database.TopRawMaxHp(list)
	panicErr(err)
	printPokes(resultList, "RAW HP:")
	resultList, err = database.TopRawMaxAtk(list)
	panicErr(err)
	printPokes(resultList, "RAW ATK:")
	resultList, err = database.TopRawMaxDef(list)
	panicErr(err)
	printPokes(resultList, "RAW DEF:")
	resultList, err = database.TopRawMaxSpAtk(list)
	panicErr(err)
	printPokes(resultList, "RAW SP.ATK:")
	resultList, err = database.TopRawMaxSpAtk(list)
	panicErr(err)
	printPokes(resultList, "RAW SP.ATK:")
	resultList, err = database.TopRawMaxSpeed(list)
	panicErr(err)
	printPokes(resultList, "RAW SPEED:")

	resultList, err = database.TopBalancedMaxHp(list)
	panicErr(err)
	printPokes(resultList, "BALANCED HP:")
	resultList, err = database.TopBalancedMaxAtk(list)
	panicErr(err)
	printPokes(resultList, "BALANCED ATK:")
	resultList, err = database.TopBalancedMaxDef(list)
	panicErr(err)
	printPokes(resultList, "BALANCED DEF:")
	resultList, err = database.TopBalancedMaxSpAtk(list)
	panicErr(err)
	printPokes(resultList, "BALANCED SP.ATK:")
	resultList, err = database.TopBalancedMaxSpDef(list)
	panicErr(err)
	printPokes(resultList, "BALANCED SP.DEF:")
	resultList, err = database.TopBalancedMaxSpeed(list)
	panicErr(err)
	printPokes(resultList, "BALANCED SPEED:")
}

func printPokes(resultList []types.Pokemon, cat string) {
	fmt.Println(cat)
	for _, pkmn := range resultList {
		fmt.Println(
			pkmn.Name,
			pkmn.Hp,
			pkmn.Atk,
			pkmn.Def,
			pkmn.SpAtk,
			pkmn.SpDef,
			pkmn.Speed,
			pkmn.Total,
		)
	}
	fmt.Println("")
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
