package main

import (
	"fmt"
	"os"

	"WhisperVault/vault"
	"WhisperVault/utils"
)

func main() {
	fmt.Println("🧙‍♂️ Welcome to WhisperVault")

	//password
	password := utils.PromptPassword("Enter vault password: ")

	//ecrypt/load contacts
	contacts, err := vault.LoadContacts("data/contacts.enc", password)
	if err != nil {
		fmt.Println("⚠️ Failed to unlock vault:", err)
		os.Exit(1)
	}

	//Add, View, Delete
	fmt.Printf("Vault unlocked. Loaded %d contact(s).\n", len(contacts))
}
