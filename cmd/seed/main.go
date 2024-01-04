package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jordipt9/gopokemon/database"
	"github.com/jordipt9/gopokemon/types"
)

const (
	nameIdx            = 1
	type1Idx           = 3
	type2Idx           = 4
	totalIdx           = 8
	hpIdx              = 9
	atkIdx             = 10
	defIdx             = 11
	spAtkIdx           = 12
	spDefIdx           = 13
	speedIdx           = 14
	againstNormalIdx   = 16
	againstFireIdx     = 17
	againstWaterIdx    = 18
	againstElectricIdx = 19
	againstGrassIdx    = 20
	againstIceIdx      = 21
	againstFightIdx    = 22
	againstPoisonIdx   = 23
	againstGroundIdx   = 24
	againstFlyingIdx   = 25
	againstPsychicIdx  = 26
	againstBugIdx      = 27
	againstRockIdx     = 28
	againstGhostIdx    = 29
	againstDragonIdx   = 30
	againstDarkIdx     = 31
	againstSteelIdx    = 32
	againstFairyIdx    = 33
)

func main() {
	database.Init()
	database.SetUpSchema()

	file, err := os.Open("cmd/seed/pokemon.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	pokemons := mapToPokemonList(records)

	err = database.InsertPokemons(pokemons)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database filled with data successfully.")
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
			AgainstNormal:   toFloat32(record[againstNormalIdx]),
			AgainstFire:     toFloat32(record[againstFireIdx]),
			AgainstWater:    toFloat32(record[againstWaterIdx]),
			AgainstElectric: toFloat32(record[againstElectricIdx]),
			AgainstGrass:    toFloat32(record[againstGrassIdx]),
			AgainstIce:      toFloat32(record[againstIceIdx]),
			AgainstFight:    toFloat32(record[againstFightIdx]),
			AgainstPoison:   toFloat32(record[againstPoisonIdx]),
			AgainstGround:   toFloat32(record[againstGroundIdx]),
			AgainstFlying:   toFloat32(record[againstFlyingIdx]),
			AgainstPsychic:  toFloat32(record[againstPsychicIdx]),
			AgainstBug:      toFloat32(record[againstBugIdx]),
			AgainstRock:     toFloat32(record[againstRockIdx]),
			AgainstGhost:    toFloat32(record[againstGhostIdx]),
			AgainstDragon:   toFloat32(record[againstDragonIdx]),
			AgainstDark:     toFloat32(record[againstDarkIdx]),
			AgainstSteel:    toFloat32(record[againstSteelIdx]),
			AgainstFairy:    toFloat32(record[againstFairyIdx]),
		}
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

func toFloat32(s string) float32 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return float32(i)
}
