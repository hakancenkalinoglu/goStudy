package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowTodoList() {
	var tasks []string

	fmt.Println("\n=== Interactive To-Do List ===")
	fmt.Println("Enter your tasks. Type 'exit' to stop and see the final list.")
	fmt.Println("-------------------------------------------------------------")

	// Boşluklu cümleleri okuyabilmek için özel bir okuyucu (Scanner) oluşturuyoruz
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a new task: ")
		scanner.Scan()
		newTask := scanner.Text()

		if strings.ToLower(newTask) == "exit" {
			break
		}

		tasks = append(tasks, newTask)
		fmt.Println(" -> Tasks added!")
	}

	fmt.Println("\n=== My Final To-Do List ===")

	if len(tasks) == 0 {
		fmt.Println("Your to-do  list is empty")
	} else {
		for index, task := range tasks {
			fmt.Printf("%d. %s", index+1, task)
		}
	}
	fmt.Println("---------------------------")

}
