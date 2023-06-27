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
type player struct {
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

//  - Print out the statistic change within each function
//  - Execute each function at least once
func main() {

}
