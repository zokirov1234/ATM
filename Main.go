package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	option := "-1"

	fmt.Println("\t\tWelcome to ATM")

	for option != "0" {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Print("Enter operation: ")
		scanner.Scan()
		option = scanner.Text()
		switch option {
		case "1":
			fmt.Println("Login")
		case "2":
			fmt.Println("Register")
		case "0":
			fmt.Println("Goodbye :)")
		default:
			fmt.Println("Wrong option")
		}
	}

}
