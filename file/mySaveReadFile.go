package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createFile() {
	_, err := os.Stat("User.txt")
	if os.IsNotExist(err) {
		file, err := os.Create("User.txt")
		if err != nil {
			fmt.Println("User.txt not created!")
			return
		}
		defer file.Close()
	}
}

func checkUserExists(email string) bool {
	file, err := os.Open("User.txt")
	if err != nil {
		fmt.Println("Error opening User.txt file.")
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) >= 3 && parts[2] == email {
			return true
		}
	}
	return false
}

func loginUser(email, password string) bool {
	file, err := os.Open("User.txt")
	if err != nil {
		fmt.Println("Error opening User.txt file.")
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) >= 4 && parts[2] == email && parts[3] == password {
			return true
		}
	}
	return false
}

func saveUser(firstName, lastName, email, password string) {
	file, err := os.OpenFile("User.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening User.txt file.")
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", firstName, lastName, email, password))
	if err != nil {
		fmt.Println("Error writing to User.txt.")
	}
}

func Run() {
	createFile()

	var firstName, lastName, email, password string
	var account int

	fmt.Println("1-Login")
	fmt.Println("2-SignUp")
	fmt.Scan(&account)

	if account == 1 {
		fmt.Println("Email: ")
		fmt.Scan(&email)
		fmt.Println("Password: ")
		fmt.Scan(&password)

		if loginUser(email, password) {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Invalid email or password!")
		}

	} else if account == 2 {
		fmt.Println("FirstName: ")
		fmt.Scan(&firstName)
		fmt.Println("LastName: ")
		fmt.Scan(&lastName)
		fmt.Println("Email: ")
		fmt.Scan(&email)
		fmt.Println("Password: ")
		fmt.Scan(&password)

		if checkUserExists(email) {
			fmt.Println("This email is already registered!")
		} else {
			saveUser(firstName, lastName, email, password)
			fmt.Println("SignUp successful!")
		}

	} else {
		fmt.Println("Invalid input! Enter 1-Login, 2-SignUp.")
	}
}
