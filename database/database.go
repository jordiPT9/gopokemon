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

func TopRawMaxHp(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("hp desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
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

func TopRawMaxDef(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("def desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func TopRawMaxSpAtk(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("sp_atk desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func TopRawMaxSpDef(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("sp_def desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func TopRawMaxSpeed(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}
	result := db.Where("LOWER(name) IN ?", names).Order("speed desc").Limit(10).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func TopBalancedMaxHp(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxHp(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
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

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxAtk(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func TopBalancedMaxDef(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxDef(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func TopBalancedMaxSpAtk(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxSpAtk(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func TopBalancedMaxSpDef(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxSpDef(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func TopBalancedMaxSpeed(names []string) ([]types.Pokemon, error) {
	var pokemons []types.Pokemon
	for i, name := range names {
		names[i] = strings.ToLower(name)
	}

	result := db.Where("LOWER(name) IN ?", names).Find(&pokemons)
	if result.Error != nil {
		return nil, result.Error
	}

	filteredPokemons := make([]types.Pokemon, 0)
	for _, pokemon := range pokemons {
		if isMaxSpeed(pokemon) {
			filteredPokemons = append(filteredPokemons, pokemon)
		}
	}

	sort.Slice(filteredPokemons, func(i, j int) bool {
		return filteredPokemons[i].Total > filteredPokemons[j].Total
	})

	if len(filteredPokemons) > 10 {
		filteredPokemons = filteredPokemons[:10]
	}

	return filteredPokemons, nil
}

func isMaxAtk(pokemon types.Pokemon) bool {
	return pokemon.Atk > pokemon.Hp &&
		pokemon.Atk > pokemon.Def &&
		pokemon.Atk > pokemon.SpAtk &&
		pokemon.Atk > pokemon.SpDef &&
		pokemon.Atk > pokemon.Speed
}

func isMaxHp(pokemon types.Pokemon) bool {
	return pokemon.Hp > pokemon.Atk &&
		pokemon.Hp > pokemon.Def &&
		pokemon.Hp > pokemon.SpAtk &&
		pokemon.Hp > pokemon.SpDef &&
		pokemon.Hp > pokemon.Speed
}

func isMaxDef(pokemon types.Pokemon) bool {
	return pokemon.Def > pokemon.Hp &&
		pokemon.Def > pokemon.Atk &&
		pokemon.Def > pokemon.SpAtk &&
		pokemon.Def > pokemon.SpDef &&
		pokemon.Def > pokemon.Speed
}

func isMaxSpAtk(pokemon types.Pokemon) bool {
	return pokemon.SpAtk > pokemon.Hp &&
		pokemon.SpAtk > pokemon.Atk &&
		pokemon.SpAtk > pokemon.Def &&
		pokemon.SpAtk > pokemon.SpDef &&
		pokemon.SpAtk > pokemon.Speed
}

func isMaxSpDef(pokemon types.Pokemon) bool {
	return pokemon.SpDef > pokemon.Hp &&
		pokemon.SpDef > pokemon.Atk &&
		pokemon.SpDef > pokemon.Def &&
		pokemon.SpDef > pokemon.SpAtk &&
		pokemon.SpDef > pokemon.Speed
}

func isMaxSpeed(pokemon types.Pokemon) bool {
	return pokemon.Speed > pokemon.Hp &&
		pokemon.Speed > pokemon.Atk &&
		pokemon.Speed > pokemon.Def &&
		pokemon.Speed > pokemon.SpAtk &&
		pokemon.Speed > pokemon.SpDef
}
