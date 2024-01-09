package database

import (
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

func InsertAllPokemon(pokemons []types.Pokemon) error {
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

func FindAll(names []string) ([]types.Pokemon, error) {
	var pokemonList []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	// result := db.Where("LOWER(name) IN ?", names).Order("total desc").Find(&pokemonList)
	result := db.Where("LOWER(name) IN ?", names).Find(&pokemonList)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemonList, nil
}

func TopBaseStats(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("total desc").Limit(20).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func percentageMore(newValue float64, originalValue float64) float64 {
	return ((newValue - originalValue) / originalValue) * 100
}

func FindTopMaxStat(names []string, stat string, maxResults int) (types.PokemonList, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order(stat + " desc").Limit(maxResults).Find(&pokemons)
	if result.Error != nil {
		return types.PokemonList{}, result.Error
	}

	return types.PokemonList{List: pokemons}, nil
}
