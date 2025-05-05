package ui

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"

	"WhisperVault/internal/models"
)

// AskPassword prompts the user for their vault password.
func AskPassword(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	password, _ := reader.ReadString('\n')
	return strings.TrimSpace(password)
}

// ShowMainMenu displays the main options and reads the user’s choice.
func ShowMainMenue() int {
	fmt.Println("\n--- WhisperVault Menu ---")
	fmt.Println("1. List Contacts")
	fmt.Println("2. Add New Contact")
	fmt.Println("3. Save and Exit")
	fmt.Print("Enter choice: ")

	var choice int
	fmt.Scanln(&choice)
	return choice
}

// List_Contacts displays all contacts in the vault.
func List_Contacts(contacts []models.Contact) {
	if len(contacts) == 0 {
		fmt.Println("No contacts found in vault.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for i, c := range contacts {
		fmt.Printf("%d. %s - %s | %s\n", i+1, c.Name, c.Phone, c.Email)
		if c.Note != "" {
			fmt.Printf("   Note: %s\n", c.Note)
		}
		fmt.Printf("   Created: %s\n", c.CreatedAt)
	}
}

// Add_Contacts prompts the user to enter new contact details.
func Add_Contacts(existing []models.Contact) []models.Contact {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Enter Email (optional): ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter Note (optional): ")
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)

	contact := models.Contact{
		Name:      name,
		Phone:     phone,
		Email:     email,
		Note:      note,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	fmt.Println("Contact added.")
	return append(existing, contact)
}
