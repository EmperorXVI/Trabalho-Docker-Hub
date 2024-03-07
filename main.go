package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     string `json:"age"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var People []Person
var File *os.File
var Content []byte
var err error

func main() {
	for {
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdString = strings.TrimSpace(cmdString)

		switch cmdString {
		case "1":
			fmt.Println("Enter Person Details")
			addpeople()
		case "2":
			fmt.Println("Get People")
			getpeople()
		case "3":
			fmt.Println("Delete Person")
			deleteperson()
		case "4":
			fmt.Println("Update Person")
			updateperson()
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func showMenu() {
	fmt.Println("1. Add Person")
	fmt.Println("2. Get People")
	fmt.Println("3. Delete Person")
	fmt.Println("4. Update Person")
	fmt.Println("5. Exit")
}

func addpeople() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("Enter Address: ")
	address, _ := reader.ReadString('\n')

	fmt.Println("Enter Age: ")
	age, _ := reader.ReadString('\n')

	fmt.Println("Enter Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Println("Enter Phone: ")
	phone, _ := reader.ReadString('\n')

	person := Person{
		Name:    strings.TrimSpace(name),
		Address: strings.TrimSpace(address),
		Age:     strings.TrimSpace(age),
		Email:   strings.TrimSpace(email),
		Phone:   strings.TrimSpace(phone),
	}

	People = loadpeople()
	People = append(People, person)
	savepeoples()
}

func loadpeople() []Person {
	file, err := os.ReadFile("people.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	_ = json.Unmarshal(file, &People)
	return People
}

func savepeoples() {
	file, _ := json.MarshalIndent(People, "", " ")
	_ = os.WriteFile("people.json", file, 0644)
}

func getpeople() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name to Search: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, person := range People {
		if person.Name == name {
			fmt.Println("Person Found:", person)
			return
		}
	}
	fmt.Println("Person Not Found")
}

func deleteperson() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Name to Delete: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	index := findPersonIndexByName(name)
	if index == -1 {
		fmt.Println("Person Not Found")
		return
	}

	People = append(People[:index], People[index+1:]...)
	savepeoples()
	fmt.Println("Person Deleted Successfully")

}

func findPersonIndexByName(name string) int {
	for i, person := range People {
		if person.Name == name {
			return i
		}
	}
	return -1
}

func updateperson() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the person's name to update: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	index := findPersonIndexByName(name)
	if index == -1 {
		fmt.Println("Person not found")
		return
	}

	fmt.Println("What information do you want to edit?")
	fmt.Println("1. Name")
	fmt.Println("2. Address")
	fmt.Println("3. Age")
	fmt.Println("4. Email")
	fmt.Println("5. Phone")

	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		updatePersonName(index, reader)
	case 2:
		updatePersonAddress(index, reader)
	case 3:
		updatePersonAge(index, reader)
	case 4:
		updatePersonEmail(index, reader)
	case 5:
		updatePersonPhone(index, reader)
	default:
		fmt.Println("Invalid option")
		return
	}

	savepeoples()
	fmt.Println("Person updated successfully")
}

func updatePersonName(index int, reader *bufio.Reader) {
	fmt.Println("Enter new name: ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	People[index].Name = newName
}

func updatePersonAddress(index int, reader *bufio.Reader) {
	fmt.Println("Enter nem Address: ")
	newAddress, _ := reader.ReadString('\n')
	newAddress = strings.TrimSpace(newAddress)

	if newAddress == "" {
		fmt.Println("address cannot be empty")
		return
	}

	People[index].Address = newAddress
}

func updatePersonAge(index int, reader *bufio.Reader) {
	fmt.Println("Enter the new age: ")
	var newAge int
	fmt.Scanf("%d", &newAge)

	if newAge < 0 || newAge > 120 {
		fmt.Println("Invalid age. Enter an age between 0 and 120.")
		return
	}

	People[index].Age = strconv.Itoa(newAge)
}

func updatePersonEmail(index int, reader *bufio.Reader) {
	fmt.Println("Enter the new email: ")
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	if newEmail == "" {
		fmt.Println("Email cannot be empty")
		return
	}

	People[index].Email = newEmail
}

func updatePersonPhone(index int, reader *bufio.Reader) {
	fmt.Println("Enter the new phone number: ")
	newPhone, _ := reader.ReadString('\n')
	newPhone = strings.TrimSpace(newPhone)

	if newPhone == "" {
		fmt.Println("Phone number cannot be empty")
		return
	}

	People[index].Phone = newPhone
}
