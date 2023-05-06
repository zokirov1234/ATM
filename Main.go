package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var users []user

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	option := -1

	fmt.Println("\t\t\t\tWelcome to ATM")

	for option != 0 {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Print("Enter operation: ")
		option = getIntInput(scanner)

		switch option {
		case 1:
			fmt.Println("Login")
		case 2:
			fmt.Println("Register")
			doRegister(scanner)
		case 0:
			fmt.Println("Goodbye :)")
		default:
			fmt.Println("Wrong option")
		}
	}

}

type user struct {
	username    string
	phoneNumber string
	password    string
	secretKey   string
}

func getIntInput(scanner *bufio.Scanner) int {
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		option, err := strconv.Atoi(input)
		if err == nil {
			return option
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}

func getStringInput(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if len(input) <= 0 {
		return "", errors.New("the input must be greater than 0")
	}
	return input, nil
}

func doRegister(scanner *bufio.Scanner) {
	fmt.Println("\t\tRegister")
	fmt.Print("Enter your name: ")
	name, err := getStringInput(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter your phone number: ")
	phoneNumber, err := getStringInput(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter your password: ")
	password, err := getStringInput(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter your secret key: ")
	secretKey, err := getStringInput(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}

	users = append(users, user{name, phoneNumber, password, secretKey})
	for _, user1 := range users {
		fmt.Println(user1)
	}
}
