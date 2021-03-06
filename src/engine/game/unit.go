package game

import (
	"api"
	"errors"
	"github.com/Tactique/golib/logger"
)

type Unit struct {
	name      string
	id        int
	health    int
	maxHealth int
	nation    nation
	movement  *Movement
	canMove   bool
	attacks   []*attack
	armor     *armor
	canAttack bool
}

func NewUnit(name string, id int, nation nation, health int, attacks []*attack, armor *armor, movement *Movement) *Unit {
	return &Unit{
		name: name, id: id, health: health, maxHealth: health,
		nation:   nation,
		movement: movement, canMove: true,
		attacks: attacks, canAttack: true,
		armor: armor}
}

func (unit *Unit) Serialize(loc Location) *api.UnitStruct {
	attacks := make([]*api.AttackStruct, len(unit.attacks))
	for i, attack := range unit.attacks {
		attacks[i] = attack.serialize()
	}
	return &api.UnitStruct{
		Name:      unit.name,
		Health:    unit.health,
		MaxHealth: unit.maxHealth,
		Nation:    int(unit.nation),
		Movement:  unit.movement.serialize(),
		Position:  loc.serialize(),
		CanMove:   unit.canMove,
		CanAttack: unit.canAttack,
		Attacks:   attacks,
		Armor:     unit.armor.serialize()}
}

func (unit *Unit) GetId() int {
	return unit.id
}

func (unit *Unit) GetMovement() *Movement {
	return unit.movement
}

func (unit *Unit) turnReset() {
	unit.canMove = true
	unit.canAttack = true
}

func (unit *Unit) receiveDamage(delta int) (bool, error) {
	logger.Debug("%s taking damage", unit.name)
	if delta < 0 {
		message := "Cannot receive a negative amount of damage"
		logger.Infof(message+" (got %d)", delta)
		return true, errors.New(message)
	}
	return unit.changeHealth(unit.health - delta), nil
}

func (unit *Unit) changeHealth(newHealth int) bool {
	logger.Infof("Unit %s is being attacked with %d health", unit.name, unit.health)
	unit.health = newHealth
	if unit.health > unit.maxHealth {
		logger.Infof("Unit %s cannot go above max health, capping from %d", unit.name, unit.health)
		unit.health = unit.maxHealth
	}
	if unit.health <= 0 {
		logger.Infof("Unit %s cannot go below 0 health, raising from from %d (unit is dead)", unit.name, unit.health)
		unit.health = 0
		return false
	}
	logger.Infof("Unit %s is now at %d health", unit.name, unit.health)
	return true
}
