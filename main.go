package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func getPasswordOrPanic(prompt string) []byte {

	fmt.Printf(prompt)

	password, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println("Could not register password")
		os.Exit(1)
	}
	return password
}

func main() {

	password := getPasswordOrPanic("Password: ")
	confirm := getPasswordOrPanic("Confirm password: ")

	if string(password) != string(confirm) {
		fmt.Println("Passwords did not match")
		os.Exit(1)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error creating password", err)
		os.Exit(1)
	}

	fmt.Println(string(hashedPassword))
}
