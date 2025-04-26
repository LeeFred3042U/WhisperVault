package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

const filename = "contacts.json"

func main() {
	for {
		fmt.Println("\n-----------ContactVault Lite-----------")
		fmt.Println("1. List Contacts")
		fmt.Println("2. Add Contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. Exit")
		fmt.Print("Choose Option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			listContacts()
		case 2:
			addContact()
		case 3:
			deleteContact()
		case 4:
			fmt.Println("Bye Maybe")
			return
		default:
			fmt.Println("Invalid Option  (^;^)")
		}
	}
}

func loadContacts() []Contact {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Contact{}
	}
	var contacts []Contact
	json.Unmarshal(data, &contacts)
	return contacts
}

func saveContacts(contacts []Contact) {
	data, _ := json.MarshalIndent(contacts, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func addContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name >;) : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	contacts := loadContacts()
	newID := 1
	if len(contacts) > 0 {
		newID = contacts[len(contacts)-1].ID + 1
	}

	contact := Contact{ID: newID, Name: name, Phone: phone}
	contacts = append(contacts, contact)
	saveContacts(contacts)

	fmt.Println("Contact Added O_o")
}

func listContacts() {
	contacts := loadContacts()
	if len(contacts) == 0 {
		fmt.Println("No Contacts Found :=(")
		return
	}

	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Phone)
	}
}

func deleteContact() {
	fmt.Print("Enter ID to delete: ")
	var id int
	fmt.Scanln(&id)

	contacts := loadContacts()
	newList := []Contact{}
	found := false

	for _, c := range contacts {
		if c.ID != id {
			newList = append(newList, c)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Contact not Found :[")
		return
	}

	saveContacts(newList)
	fmt.Println("Contact Deleted :)")
}
