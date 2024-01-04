package types

import "fmt"

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
	AgainstNormal   float32 `gorm:"column:against_normal"`
	AgainstFire     float32 `gorm:"column:against_fire"`
	AgainstWater    float32 `gorm:"column:against_water"`
	AgainstElectric float32 `gorm:"column:against_electric"`
	AgainstGrass    float32 `gorm:"column:against_grass"`
	AgainstIce      float32 `gorm:"column:against_ice"`
	AgainstFight    float32 `gorm:"column:against_fight"`
	AgainstPoison   float32 `gorm:"column:against_poison"`
	AgainstGround   float32 `gorm:"column:against_ground"`
	AgainstFlying   float32 `gorm:"column:against_flying"`
	AgainstPsychic  float32 `gorm:"column:against_psychic"`
	AgainstBug      float32 `gorm:"column:against_bug"`
	AgainstRock     float32 `gorm:"column:against_rock"`
	AgainstGhost    float32 `gorm:"column:against_ghost"`
	AgainstDragon   float32 `gorm:"column:against_dragon"`
	AgainstDark     float32 `gorm:"column:against_dark"`
	AgainstSteel    float32 `gorm:"column:against_steel"`
	AgainstFairy    float32 `gorm:"column:against_fairy"`
}

func (p Pokemon) String() string {
	return fmt.Sprintf(`
		Name: %s
		Type1: %s
		Type2: %s
		Hp: %d
		Atk: %d
		Def: %d
		SpAtk: %d
		SpDef: %d
		Speed: %d
		Total: %d
		`,
		p.Name,
		p.Type1,
		p.Type2,
		p.Hp,
		p.Atk,
		p.Def,
		p.SpAtk,
		p.SpDef,
		p.Speed,
		p.Total,
	)
}
