package main

import "fmt"

//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
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
	fmt.Println(player.name, "Consumed", amount, "energy ->", player.health)
}

//  - Print out the statistic change within each function
//  - Execute each function at least once
func main() {

}
