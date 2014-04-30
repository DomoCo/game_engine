package game_engine

import (
	"api"
)

type attackType int

const (
	sword attackType = iota
	axe
	mace
	wand
	staff
)

type attack struct {
	name       string
	attackType attackType
	power      int
	minRange   uint
	maxRange   uint
}

func newAttack(name string, attackType attackType, power int, minRange uint, maxRange uint) *attack {
	return &attack{
		name:       name,
		attackType: attackType,
		power:      power,
		minRange:   minRange,
		maxRange:   maxRange,
	}
}

func (attack *attack) serialize() *api.AttackStruct {
	return &api.AttackStruct{
		Name:       attack.name,
		AttackType: int(attack.attackType),
		Power:      attack.power,
		MinRange:   attack.minRange,
		MaxRange:   attack.maxRange,
	}
}
