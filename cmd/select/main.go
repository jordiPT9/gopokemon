package main

import (
	"fmt"
	"os"

	"github.com/jordipt9/gopokemon/database"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Usage: make select name=<pokemon_name>")
		return
	}

	database.Init()
	pkmnName := os.Args[1]
	pkmn, err := database.FindByName(pkmnName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pkmn)
	}
}
