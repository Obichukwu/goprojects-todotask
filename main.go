/*
Copyright © 2024 Onigns <keita.onigns@outlook.com>
*/
package main

import (
	"onigns.io/keita/todotasks/cmd"
	"onigns.io/keita/todotasks/models"
)

func main() {
	models.LoadFromFile("tasks.json")

	cmd.Execute()

	models.SaveTaskManagerToFile("tasks.json")
}
