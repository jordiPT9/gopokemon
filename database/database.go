package database

import (
	"sort"
	"strings"

	"github.com/jordipt9/gopokemon/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	database, err := gorm.Open(sqlite.Open("pokedex.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

func SetUpSchema() {
	err := db.AutoMigrate(&types.Pokemon{})
	if err != nil {
		panic(err)
	}
}

func DropPokemonsTable() {
	if err := db.Exec("DELETE FROM pokemons").Error; err != nil {
		panic(err)
	}
}

func InsertPokemons(pokemons []types.Pokemon) error {
	result := db.CreateInBatches(pokemons, len(pokemons))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindByName(name string) (types.Pokemon, error) {
	var pkmn types.Pokemon
	result := db.Where("LOWER(name) = ?", strings.ToLower(name)).First(&pkmn)
	if result.Error != nil {
		return types.Pokemon{}, result.Error
	}
	return pkmn, nil
}

func TopRawMaxAtk(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("atk desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func TopBalancedMaxAtk(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	// Filtrar por 'atk' más alto
	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxAtk(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	// Ordenar por 'total'
	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	// Tomar solo los primeros 10
	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func isMaxAtk(pokemon types.Pokemon) bool {
	return pokemon.Atk >= pokemon.Hp &&
		pokemon.Atk >= pokemon.Def &&
		pokemon.Atk >= pokemon.SpAtk &&
		pokemon.Atk >= pokemon.SpDef &&
		pokemon.Atk >= pokemon.Speed
}
