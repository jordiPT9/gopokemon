package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jordipt9/gopokemon/database"
	"github.com/jordipt9/gopokemon/types"
)

// const (
// 	nameIdx            = 1
// 	type1Idx           = 3
// 	type2Idx           = 4
// 	totalIdx           = 8
// 	hpIdx              = 9
// 	atkIdx             = 10
// 	defIdx             = 11
// 	spAtkIdx           = 12
// 	spDefIdx           = 13
// 	speedIdx           = 14
// 	againstNormalIdx   = 16
// 	againstFireIdx     = 17
// 	againstWaterIdx    = 18
// 	againstElectricIdx = 19
// 	againstGrassIdx    = 20
// 	againstIceIdx      = 21
// 	againstFightIdx    = 22
// 	againstPoisonIdx   = 23
// 	againstGroundIdx   = 24
// 	againstFlyingIdx   = 25
// 	againstPsychicIdx  = 26
// 	againstBugIdx      = 27
// 	againstRockIdx     = 28
// 	againstGhostIdx    = 29
// 	againstDragonIdx   = 30
// 	againstDarkIdx     = 31
// 	againstSteelIdx    = 32
// 	againstFairyIdx    = 33
// )

const (
	nameIdx            = 2
	type1Idx           = 9
	type2Idx           = 10
	totalIdx           = 17
	hpIdx              = 18
	atkIdx             = 19
	defIdx             = 20
	spAtkIdx           = 21
	spDefIdx           = 22
	speedIdx           = 23
	againstNormalIdx   = 33
	againstFireIdx     = 34
	againstWaterIdx    = 35
	againstElectricIdx = 36
	againstGrassIdx    = 37
	againstIceIdx      = 38
	againstFightIdx    = 39
	againstPoisonIdx   = 40
	againstGroundIdx   = 41
	againstFlyingIdx   = 42
	againstPsychicIdx  = 43
	againstBugIdx      = 44
	againstRockIdx     = 45
	againstGhostIdx    = 46
	againstDragonIdx   = 47
	againstDarkIdx     = 48
	againstSteelIdx    = 49
	againstFairyIdx    = 50
)

func main() {
	srcFile := "pokemon_gen_8.csv"
	if len(os.Args) == 2 {
		srcFile = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("Usage: make select name=<pokemon_name>")
		return
	}

	database.Init()
	database.SetUpSchema()

	path := fmt.Sprintf("cmd/seed/%s", srcFile)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	pokemonList := mapToPokemonList(records)

	err = database.InsertAllPokemon(pokemonList)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database filled with data successfully with source file: ", srcFile)
}

func mapToPokemonList(records [][]string) []types.Pokemon {
	pokemons := []types.Pokemon{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		newPkmn := types.Pokemon{
			Name:            record[nameIdx],
			Type1:           record[type1Idx],
			Type2:           record[type2Idx],
			Hp:              toInt(record[hpIdx]),
			Atk:             toInt(record[atkIdx]),
			Def:             toInt(record[defIdx]),
			SpAtk:           toInt(record[spAtkIdx]),
			SpDef:           toInt(record[spDefIdx]),
			Speed:           toInt(record[speedIdx]),
			Total:           toInt(record[totalIdx]),
			AgainstNormal:   tofloat64(record[againstNormalIdx]),
			AgainstFire:     tofloat64(record[againstFireIdx]),
			AgainstWater:    tofloat64(record[againstWaterIdx]),
			AgainstElectric: tofloat64(record[againstElectricIdx]),
			AgainstGrass:    tofloat64(record[againstGrassIdx]),
			AgainstIce:      tofloat64(record[againstIceIdx]),
			AgainstFight:    tofloat64(record[againstFightIdx]),
			AgainstPoison:   tofloat64(record[againstPoisonIdx]),
			AgainstGround:   tofloat64(record[againstGroundIdx]),
			AgainstFlying:   tofloat64(record[againstFlyingIdx]),
			AgainstPsychic:  tofloat64(record[againstPsychicIdx]),
			AgainstBug:      tofloat64(record[againstBugIdx]),
			AgainstRock:     tofloat64(record[againstRockIdx]),
			AgainstGhost:    tofloat64(record[againstGhostIdx]),
			AgainstDragon:   tofloat64(record[againstDragonIdx]),
			AgainstDark:     tofloat64(record[againstDarkIdx]),
			AgainstSteel:    tofloat64(record[againstSteelIdx]),
			AgainstFairy:    tofloat64(record[againstFairyIdx]),
		}
		totalWeaknesses := newPkmn.AgainstNormal +
			newPkmn.AgainstFire +
			newPkmn.AgainstWater +
			newPkmn.AgainstElectric +
			newPkmn.AgainstGrass +
			newPkmn.AgainstIce +
			newPkmn.AgainstFight +
			newPkmn.AgainstPoison +
			newPkmn.AgainstGround +
			newPkmn.AgainstFlying +
			newPkmn.AgainstPsychic +
			newPkmn.AgainstBug +
			newPkmn.AgainstRock +
			newPkmn.AgainstGhost +
			newPkmn.AgainstDragon +
			newPkmn.AgainstDark +
			newPkmn.AgainstSteel +
			newPkmn.AgainstFairy

		newPkmn.TotalWeaknesses = totalWeaknesses
		pokemons = append(pokemons, newPkmn)
	}
	return pokemons
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func tofloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return float64(i)
}
