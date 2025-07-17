package main

import (
	"fmt"

	"time"

	"github.com/ganeswar-velvadapu/file-manager/operations"
)

func main() {

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - List all files and folders in current working directory")
		fmt.Println("2 - Open and read a file")
		fmt.Println("3 - Create a file")
		fmt.Println("4 - Edit a file")
		fmt.Println("5 - Delete a file")
		fmt.Println("6 - Create a directory")
		fmt.Println("7 - You wanted to move into a directory")
		fmt.Println("0 - Exit")

		fmt.Print("Enter your choice: ")
		var i int
		_, err := fmt.Scanln(&i)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch i {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			fmt.Println("You wanted to list all files and folders in current working directory")
			operations.ListAll()
		case 2:
			fmt.Println("You wanted to open and read a file")
			operations.OpenAndReadFile()
		case 3:
			fmt.Println("You wanted to create a file")
			time.Sleep(1 * time.Second)
		case 4:
			fmt.Println("You wanted to edit a file")
			time.Sleep(1 * time.Second)
		case 5:
			fmt.Println("You wanted to delete a file")
			time.Sleep(1 * time.Second)
		case 6:
			fmt.Println("You want to create a directory inside current working directory")
			time.Sleep(1 * time.Second)
		case 7:
			fmt.Println("You want to move into a directory")
			time.Sleep(1 * time.Second)
		default:
			fmt.Println("Please select a valid option")
		}
	}
}
