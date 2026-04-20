package main

import "fmt"

func main() {
	var choice string

	fmt.Println("\n=== LESSON 1: MINI PROJECTS ===")
	fmt.Println("1. Professional Calculator")
	fmt.Println("2. Number Guessing Game")
	fmt.Println("3. My To-Do List (Slices & Loops)") // Yeni seçeneğimiz
	fmt.Print("Which program do you want to run? (Press 1, 2, or 3): ")
	fmt.Scanln(&choice)

	if choice == "1" {
		// ... (Burada daha önce yazdığın hesap makinesi kodları var, onlara dokunma)
		fmt.Println("Calculator starting...")

	} else if choice == "2" {
		PlayGuessGame()

	} else if choice == "3" { // Yeni 3. seçeneğimizin kontrolü
		ShowTodoList()

	} else {
		fmt.Println("Invalid choice. Please run the program again.")
	}
}
