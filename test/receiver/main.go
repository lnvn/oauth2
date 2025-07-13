package main

import "fmt"

type Character struct {
	Name   string
	Health int
}

func (c *Character) takeDamage(amount int) {
	fmt.Printf("-> The character '%s' is taking %d damage! \n", c.Name, amount)
	c.Health -= amount

	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) displayStatus() {
	fmt.Printf("Status:\n\tName: %s\n\tHealth: %d HP\n", c.Name, c.Health)
}

func main() {
	player := Character{Name: "Long", Health: 100}
	fmt.Println(" --- Game Start ---")
	player.displayStatus()

	fmt.Print("\n")
	fmt.Println(" --- Big Enemy ---")
	player.takeDamage(20)
	player.displayStatus()

	fmt.Print("\n")
	fmt.Println(" --- Giant Enemy ---")
	player.takeDamage(40)
	player.displayStatus()

	fmt.Print("\n")
	fmt.Println(" --- K/O ---")
	player.takeDamage(80)
	player.displayStatus()
}
