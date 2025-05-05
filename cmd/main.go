package main

import (
	"fmt"
	"log"
	"os"

	"WhisperVault/internal/models"
	"WhisperVault/internal/ui"
	"WhisperVault/internal/vault"
)

func main() {
	const vaultPath = "vault_data.enc"

	var (
		contacts []models.Contact
		password string
		err      error
	)

	// --- Authentication ---
	maxRetries := 3
	for attempts := 0; attempts < maxRetries; attempts++ {
		password = ui.AskPassword("Enter Vault Password: ")

		contacts, err = vault.LoadVault(vaultPath, password)
		if err == nil {
			fmt.Println("Success: Vault Unlocked")
			break
		}

		fmt.Println("Invalid Password or Corrupted Vault.")
		if attempts == maxRetries-1 {
			fmt.Println("Too many failed attempts. Exiting.")
			os.Exit(1)
		}
	}

	// --- Main Menu Loop ---
	for {
		choice := ui.ShowMainMenue()

		switch choice {
		case 1:
			ui.List_Contacts(contacts)
		case 2:
			contacts = ui.Add_Contacts(contacts)
		case 3:
			err := vault.SaveVault(vaultPath, contacts, password)
			if err != nil {
				log.Fatalf("Failed to save vault: %v", err)
			}
			fmt.Println("Vault Saved and Locked. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option. Try Again.")
		}
	}
}
