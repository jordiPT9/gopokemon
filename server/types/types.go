package types

import (
	"fmt"
)

const (
	Hp    string = "hp"
	Atk   string = "atk"
	Def   string = "def"
	SpAtk string = "sp_atk"
	SpDef string = "sp_def"
	Speed string = "speed"
	Total string = "total"
)

type Pokemon struct {
	Name            string  `gorm:"primaryKey"`
	Type1           string  `gorm:"type_1"`
	Type2           string  `gorm:"type_2"`
	Hp              int     `gorm:"column:hp"`
	Atk             int     `gorm:"column:atk"`
	Def             int     `gorm:"column:def"`
	SpAtk           int     `gorm:"column:sp_atk"`
	SpDef           int     `gorm:"column:sp_def"`
	Speed           int     `gorm:"column:speed"`
	Total           int     `gorm:"column:total"`
	AgainstNormal   float64 `gorm:"column:against_normal"`
	AgainstFire     float64 `gorm:"column:against_fire"`
	AgainstWater    float64 `gorm:"column:against_water"`
	AgainstElectric float64 `gorm:"column:against_electric"`
	AgainstGrass    float64 `gorm:"column:against_grass"`
	AgainstIce      float64 `gorm:"column:against_ice"`
	AgainstFight    float64 `gorm:"column:against_fight"`
	AgainstPoison   float64 `gorm:"column:against_poison"`
	AgainstGround   float64 `gorm:"column:against_ground"`
	AgainstFlying   float64 `gorm:"column:against_flying"`
	AgainstPsychic  float64 `gorm:"column:against_psychic"`
	AgainstBug      float64 `gorm:"column:against_bug"`
	AgainstRock     float64 `gorm:"column:against_rock"`
	AgainstGhost    float64 `gorm:"column:against_ghost"`
	AgainstDragon   float64 `gorm:"column:against_dragon"`
	AgainstDark     float64 `gorm:"column:against_dark"`
	AgainstSteel    float64 `gorm:"column:against_steel"`
	AgainstFairy    float64 `gorm:"column:against_fairy"`
	TotalWeaknesses float64 `gorm:"column:total_weaknesses"`
}

func (p Pokemon) GetHighestStats() string {
	stats := map[string]int{
		"HP":      p.Hp,
		"ATTACK":  p.Atk,
		"DEFENSE": p.Def,
		"SP.ATK":  p.SpAtk,
		"SP.DEF":  p.SpDef,
		"SPEED":   p.Speed,
	}

	highestValue1, highestValue2, highestValue3 := 0, 0, 0
	highestKey1, highestKey2, highestKey3 := "", "", ""

	for key, value := range stats {
		if value > highestValue1 {
			highestKey3 = highestKey2
			highestValue3 = highestValue2

			highestKey2 = highestKey1
			highestValue2 = highestValue1

			highestValue1 = value
			highestKey1 = key
		} else if value > highestValue2 {
			highestValue3 = highestValue2
			highestKey3 = highestKey2

			highestValue2 = value
			highestKey2 = key
		} else if value > highestValue3 {
			highestValue3 = value
			highestKey3 = key
		}
	}

	return fmt.Sprintf("[%s(%d)\t%s(%d)\t%s(%d)]", highestKey1, highestValue1, highestKey2, highestValue2, highestKey3, highestValue3)
}

func (p Pokemon) String() string {
	avgStatValue := float64(p.Total) / float64(6)
	var hpString string
	if p.Hp > int(avgStatValue) {
		hpString = fmt.Sprintf("*%d(Hp)", p.Hp)
	} else {
		hpString = fmt.Sprintf("%d(Hp)", p.Hp)
	}
	var atkString string
	if p.Atk > int(avgStatValue) {
		atkString = fmt.Sprintf("*%d(Atk)", p.Atk)
	} else {
		atkString = fmt.Sprintf("%d(Atk)", p.Hp)
	}
	var defString string
	if p.Def > int(avgStatValue) {
		defString = fmt.Sprintf("*%d(Def)", p.Def)
	} else {
		defString = fmt.Sprintf("%d(Def)", p.Hp)
	}
	var spAtkString string
	if p.SpAtk > int(avgStatValue) {
		spAtkString = fmt.Sprintf("*%d(SpAtk)", p.SpAtk)
	} else {
		spAtkString = fmt.Sprintf("%d(SpAtk)", p.Hp)
	}
	var spDefString string
	if p.SpDef > int(avgStatValue) {
		spDefString = fmt.Sprintf("*%d(SpDef)", p.SpDef)
	} else {
		spDefString = fmt.Sprintf("%d(SpDef)", p.Hp)
	}
	var speedString string
	if p.Speed > int(avgStatValue) {
		speedString = fmt.Sprintf("*%d(Speed)", p.Speed)
	} else {
		speedString = fmt.Sprintf("%d(Speed)", p.Hp)
	}

	return fmt.Sprintf(
		"%s \t%s \t%s \t%s \t%s \t%s \t%s \t%d(Total) \t%.2f(Weaknesses) \t%s %s",
		p.Name, hpString, atkString, defString, spAtkString, spDefString, speedString, p.Total, p.TotalWeaknesses, p.Type1, p.Type2,
	)
}

type PokemonList struct {
	List            []Pokemon
	Score           float64
	TotalStats      float64
	TotalWeaknesses float64
}

func (p PokemonList) Len() int {
	return len(p.List)
}
func (p PokemonList) Swap(i, j int) {
	p.List[i], p.List[j] = p.List[j], p.List[i]
}
func (p PokemonList) Less(i, j int) bool {
	return p.List[i].Name < p.List[j].Name
}
