package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/jordipt9/gopokemon/database"
	"github.com/jordipt9/gopokemon/types"
)

func main() {
	database.Init()

	list := []string{
		"Gothitelle",
		"Watchog",
		"vanillish",
		"cofagrigus",
		"zebstrika",
		//"maractus",
		"stoutland",
		"emboar",
		"galvantula",
		"sawk",
		"gurdurr",
		"boldore",
		"stunfisk",
		"mienfoo",
		"lampent",
		"Sawsbuck",
		"seismitoad",
		"swoobat",
		//"cinccino",
		"CHANDELURE",
		"golurk",
	}

	pokemonList, err := database.FindAll(list)
	panicErr(err)

	bestTeam, bestTeams := findBestTeam(pokemonList)
	fmt.Println("")
	fmt.Println("BEST TEAM")
	printPokemonList(bestTeam, true)

	fmt.Println("")
	fmt.Println("BEST TEAMS")
	for _, team := range bestTeams {
		printPokemonList(team, true)
	}

	//pokemons, _ := database.FindTopMaxStat(list, types.SpAtk, 20)
	//fmt.Println("")
	//fmt.Println("TOP MAX SPATK")
	//printPokemonList(pokemons, false)
	pokemons := []string{
		"Emboar",
		"zebstrika",
		"stoutland",
		"seismitoad",
		"chandelure",
		"sawsbuck",
	}
	lstss, _ := database.FindAll(pokemons)
	printPokemonList(types.PokemonList{List: lstss}, false)
}

func findBestTeam(pkmnList []types.Pokemon) (types.PokemonList, []types.PokemonList) {
	bestTeam := types.PokemonList{}
	bestTeams := []types.PokemonList{
		{},
		{},
		{},
	}
	bestScore0 := float64(0)
	bestScore1 := float64(0)
	bestScore2 := float64(0)
	counter := 0
	currentScore := float64(0)
	currentTotalStats := float64(0)
	currentTotalWeaknesses := float64(0)
	for i := 0; i < len(pkmnList)-5; i++ {
		for j := i + 1; j < len(pkmnList)-4; j++ {
			for k := j + 1; k < len(pkmnList)-3; k++ {
				for l := k + 1; l < len(pkmnList)-2; l++ {
					for m := l + 1; m < len(pkmnList)-1; m++ {
						for n := m + 1; n < len(pkmnList); n++ {
							counter += 1
							currentTeam := []types.Pokemon{pkmnList[i], pkmnList[j], pkmnList[k], pkmnList[l], pkmnList[m], pkmnList[n]}
							currentScore, currentTotalStats, currentTotalWeaknesses = calculateTeamScore(currentTeam)
							currentPokemonList := types.PokemonList{
								List:            currentTeam,
								Score:           currentScore,
								TotalStats:      currentTotalStats,
								TotalWeaknesses: currentTotalWeaknesses,
							}

							if teamHasCriticalWeakness(currentTeam, false) {
								continue
							}

							if currentScore > bestScore0 {
								bestScore0 = currentScore
								bestTeam = currentPokemonList
								bestTeams[2] = bestTeams[1]
								bestTeams[1] = bestTeams[0]
								bestTeams[0] = currentPokemonList
							} else if currentScore > bestScore1 {
								bestScore1 = currentScore
								bestTeams[2] = bestTeams[1]
								bestTeams[1] = currentPokemonList
							} else if currentScore > bestScore2 {
								bestScore2 = currentScore
								bestTeams[2] = currentPokemonList
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Nº iterations: ", counter)
	fmt.Println("Best team score: ", bestScore0)
	fmt.Println("")
	teamHasCriticalWeakness(bestTeam.List, true)
	return bestTeam, bestTeams
}

func teamHasCriticalWeakness(team []types.Pokemon, print bool) bool {
	totalAgainstNormal := []int{0, 0}
	totalAgainstFire := []int{0, 0}
	totalAgainstWater := []int{0, 0}
	totalAgainstElectric := []int{0, 0}
	totalAgainstGrass := []int{0, 0}
	totalAgainstIce := []int{0, 0}
	totalAgainstFight := []int{0, 0}
	totalAgainstPoison := []int{0, 0}
	totalAgainstGround := []int{0, 0}
	totalAgainstFlying := []int{0, 0}
	totalAgainstPsychic := []int{0, 0}
	totalAgainstBug := []int{0, 0}
	totalAgainstRock := []int{0, 0}
	totalAgainstGhost := []int{0, 0}
	totalAgainstDragon := []int{0, 0}
	totalAgainstDark := []int{0, 0}
	totalAgainstSteel := []int{0, 0}
	totalAgainstFairy := []int{0, 0}
	for _, pkmn := range team {
		totalAgainstNormal = addWeaknessScore(totalAgainstNormal, pkmn.AgainstNormal)
		totalAgainstFire = addWeaknessScore(totalAgainstFire, pkmn.AgainstFire)
		totalAgainstWater = addWeaknessScore(totalAgainstWater, pkmn.AgainstWater)
		totalAgainstElectric = addWeaknessScore(totalAgainstElectric, pkmn.AgainstElectric)
		totalAgainstGrass = addWeaknessScore(totalAgainstGrass, pkmn.AgainstGrass)
		totalAgainstIce = addWeaknessScore(totalAgainstIce, pkmn.AgainstIce)
		totalAgainstFight = addWeaknessScore(totalAgainstFight, pkmn.AgainstFight)
		totalAgainstPoison = addWeaknessScore(totalAgainstPoison, pkmn.AgainstPoison)
		totalAgainstGround = addWeaknessScore(totalAgainstGround, pkmn.AgainstGround)
		totalAgainstFlying = addWeaknessScore(totalAgainstFlying, pkmn.AgainstFlying)
		totalAgainstPsychic = addWeaknessScore(totalAgainstPsychic, pkmn.AgainstPsychic)
		totalAgainstBug = addWeaknessScore(totalAgainstBug, pkmn.AgainstBug)
		totalAgainstRock = addWeaknessScore(totalAgainstRock, pkmn.AgainstRock)
		totalAgainstGhost = addWeaknessScore(totalAgainstGhost, pkmn.AgainstGhost)
		totalAgainstDragon = addWeaknessScore(totalAgainstDragon, pkmn.AgainstDragon)
		totalAgainstDark = addWeaknessScore(totalAgainstDark, pkmn.AgainstDark)
		totalAgainstSteel = addWeaknessScore(totalAgainstSteel, pkmn.AgainstSteel)
		totalAgainstFairy = addWeaknessScore(totalAgainstFairy, pkmn.AgainstFairy)
	}
	table := map[string][]int{
		"totalAgainstNormal":   totalAgainstNormal,
		"totalAgainstFire":     totalAgainstFire,
		"totalAgainstWater":    totalAgainstWater,
		"totalAgainstElectric": totalAgainstElectric,
		"totalAgainstGrass":    totalAgainstGrass,
		"totalAgainstIce":      totalAgainstIce,
		"totalAgainstFight":    totalAgainstFight,
		"totalAgainstPoison":   totalAgainstPoison,
		"totalAgainstGround":   totalAgainstGround,
		"totalAgainstFlying":   totalAgainstFlying,
		"totalAgainstPsychic":  totalAgainstPsychic,
		"totalAgainstBug":      totalAgainstBug,
		"totalAgainstRock":     totalAgainstRock,
		"totalAgainstGhost":    totalAgainstGhost,
		"totalAgainstDragon":   totalAgainstDragon,
		"totalAgainstDark":     totalAgainstDark,
		"totalAgainstSteel":    totalAgainstSteel,
		"totalAgainstFairy":    totalAgainstFairy,
	}

	if print == true {
		fmt.Println(table)
	}

	for _, value := range table {
		if value[0] >= 3 && value[1] == 0 {
			return true
		} else if value[0] >= 4 && value[1] == 1 {
			return true
		}
	}

	return false
}

func addWeaknessScore(n []int, value float64) []int {
	switch value {
	case 0.0:
		n[1] += 1
	case 0.25:
		n[1] += 1
	case 0.5:
		n[1] += 1
	case 1.0:
		n[1] += 0
	case 2.0:
		n[0] += 1
	case 4.0:
		n[0] += 1
	}
	return n
}
func calculateTeamStats(team []types.Pokemon) int {
	totalStats := 0
	for _, pkmn := range team {
		totalStats += pkmn.Total
	}
	return totalStats
}
func calculateTeamScore(currentTeam []types.Pokemon) (float64, float64, float64) {
	teamScore := float64(0)
	totalStats := float64(0)
	totalWeaknesses := float64(0)
	for _, pokemon := range currentTeam {
		pkmnScore := float64(pokemon.Total) - float64(pokemon.TotalWeaknesses)
		teamScore += pkmnScore
		totalStats += float64(pokemon.Total)
		totalWeaknesses += pokemon.TotalWeaknesses
	}
	return teamScore, totalStats, totalWeaknesses
}

func printPokemonList(currentTeam types.PokemonList, withSorting bool) {
	sort.Sort(currentTeam)

	fmt.Println("---")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, " NAME\tHP\tATK\tDEF\tSP.ATK\tSP.DEF\tSPEED\tTOTAL\tWEAKNESSES\tTYPING\tHIGHEST")

	totalStats := 0
	totalWeaknesses := float64(0)

	for _, p := range currentTeam.List {
		avgStatValue := float64(p.Total) / float64(6)
		hpString := formatStat(p.Hp, int(avgStatValue))
		atkString := formatStat(p.Atk, int(avgStatValue))
		defString := formatStat(p.Def, int(avgStatValue))
		spAtkString := formatStat(p.SpAtk, int(avgStatValue))
		spDefString := formatStat(p.SpDef, int(avgStatValue))
		speedString := formatStat(p.Speed, int(avgStatValue))

		line := fmt.Sprintf(" %s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%.2f\t%s\t%s", p.Name, hpString, atkString, defString, spAtkString, spDefString, speedString, p.Total, p.TotalWeaknesses, p.Type1+" "+p.Type2, p.GetHighestStats())
		fmt.Fprintln(w, line)
		totalStats += p.Total
		totalWeaknesses += p.TotalWeaknesses
	}
	w.Flush()

	fmt.Println("Total stats: ", totalStats)
	fmt.Println("Total weaknesses: ", totalWeaknesses)
}

func formatStat(value, avgValue int) string {
	a := float64(value) - float64(avgValue)
	b := float64(avgValue) * 0.05
	if a > b {
		return fmt.Sprintf("%d(*)", value)
	}
	return fmt.Sprintf("%d", value)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
