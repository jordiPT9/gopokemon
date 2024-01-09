package main

import (
	"fmt"

	"github.com/jordipt9/gopokemon/database"
)

func main() {
	database.Init()
	database.DropPokemonsTable()
	fmt.Println("Database table data dropped succesfully.")
}
