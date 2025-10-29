package main

import (
	"fmt"
	"slices"
)

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

// Modifies the player's inventory
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, Item{
		Name: item.Name,
		Type: item.Type,
	})
}

// Removes the item from the player's inventory
func (p *Player) DropItem(itemName string) {
	// p.Inventory = slices.DeleteFunc(p.Inventory, func(item Item) bool {
	// 	return item.Name == itemName
	// })

	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
		}
	}
}

// Uses the item in the player's inventory
func (p *Player) UseItem(itemName string) {
	exists := slices.ContainsFunc(p.Inventory, func(item Item) bool {
		return item.Name == itemName
	})

	if exists {
		if itemName == "Sword" {
			fmt.Println("You swing the sword at the monster")
		} else if itemName == "Shield" {
			fmt.Println("You block the monster's attack")
		} else if itemName == "Bow" {
			fmt.Println("You fire an arrow at the monster")
		} else if itemName == "Potion" {
			// Potien is one time use
			fmt.Println("You drink the potion and feel refreshed")
			p.Inventory = slices.DeleteFunc(p.Inventory, func(item Item) bool {
				return item.Name == itemName
			})
		}
	} else {
		fmt.Println("Item not found")
	}
}

func main() {
	player := Player{
		Name: "John Doe",
		Inventory: []Item{
			{Name: "Sword", Type: "Weapon"},
			{Name: "Shield", Type: "Armor"},
			{Name: "Bow", Type: "Ranged"},
			{Name: "Potion", Type: "Consumable"},
		},
	}

	player.PickUpItem(Item{Name: "Axe", Type: "Weapon"})

	player.UseItem("Sword")

	player.DropItem("Sword")
	player.UseItem("Sword")
	player.UseItem("Potion")
	player.UseItem("Bow")
}
