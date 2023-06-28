package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Add", amount, "health ->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "health ->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "Add", amount, "energy ->", player.energy)
}

func (player *Player) consumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "energy ->", player.health)
}

func main() {
	player := Player{
		name:      "Scar Face",
		health:    100,
		maxHealth: 100,
		energy:    200,
		maxEnergy: 200,
	}
	player.addHealth(10)
	player.applyDamage(20)
	player.addEnergy(20)
	player.consumeEnergy(40)
}
