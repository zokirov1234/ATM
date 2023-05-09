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
		fmt.Println("\n1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Print("Enter operation: ")
		option = getIntInput(scanner)

		switch option {
		case 1:
			doLogin(scanner)
		case 2:
			doRegister(scanner)
		case 0:
			fmt.Println("Goodbye :)")
		case 3:
			for _, user := range users {
				fmt.Println(user)
			}
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
	state       bool
	cards       card
}

type card struct {
	cardNumber int
	ownerName  string
	balance    float32
	isActive   bool
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

	users = append(users, user{name, phoneNumber, password, secretKey, true, card{}})
	fmt.Println("Successfully registered")
}

func doLogin(scanner *bufio.Scanner) {
	fmt.Println("\t\tLogin")

	fmt.Print("Enter your phone number : ")
	phoneNumber, err := getStringInput(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}
	user1 := user{}
	if len(users) == 0 {
		fmt.Println("Go to register.")
		return
	}

	for index, user := range users {
		if phoneNumber == user.phoneNumber {
			user1 = user
			break
		} else {
			if index == len(users)-1 {
				fmt.Println("Wrong phone number !")
				return
			}
			continue
		}
	}
	if user1.state {
		count := 0
		for i := 0; i < 3; i++ {
			fmt.Print("Enter your password : ")
			password, err := getStringInput(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}

			if password != user1.password {
				fmt.Println("Wrong password, ", 2-count, " attempts left")
				count++
				if count == 3 {
					fmt.Println("Your account has been blocked")
					for index, user2 := range users {
						if user2.phoneNumber == phoneNumber {
							users[index].state = false
							return
						}
					}
				}
			} else {
				break
			}
		}

		userMenu(scanner, user1)

	} else {
		fmt.Println("Your account was blocked. You have to recover it.")
		return
	}
}

func userMenu(scanner *bufio.Scanner, user3 user) {

	option := -1

	for option != 0 {
		fmt.Println("\n***** Menu *****")
		fmt.Println("1. About me")
		fmt.Println("2. Balance")
		fmt.Println("3. Transfer")
		fmt.Println("4. Converter")
		fmt.Println("5. Withdraw")
		fmt.Println("5. Payments")
		fmt.Println("6. History")
		fmt.Println("7. Message to admin")
		fmt.Println("8. Save password")
		fmt.Println("9. Change password")
		fmt.Println("0. Exit")
		fmt.Print("Enter operation : ")
		option = getIntInput(scanner)

		switch option {
		case 1:
			getData(user3)
		case 2:

		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
		case 0:

		}
	}

}

func getData(user2 user) {

	fmt.Println("\n-----Me-----")

	fmt.Println("Your name : ", user2.username)
	fmt.Println("Your phone number : ", user2.phoneNumber)
	fmt.Println("Your password : ", user2.password)
	fmt.Println("Your secret key : ", user2.secretKey)
}
