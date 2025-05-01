package main

import {
    "fmt"
    "os"
    "bufio"
    "strings"
    "WhisperVault/internal/vault"
    "WhisperVault/internal/models"
    "WhisperVault/internal/ui"
}

func main() {
    const vaultPath = "vault_data.enc"

    // ----     Auhentication   ----
    var {
        contacts []models.Contact
        password string
        err      error
    }

    maxRetries := 3
    for attempts := 0; attempts < maxRetries; attempts++{
    password = ui.AskPassword("Enter Vault Password: ")

    if err == nil {
        fmt.Println("Success: Vault Unlocked")
        break
    }

    fmt.Println("Invalid Passwrod or Vault Corrupted.")
    if attempts == maxRetries-1 {
        fmt.Println("Too many failed attempts. Exiting Vault Gate")
        os.Exit(1)
    }
}       

// ----     Main Menue Loop     ----
    for {

        choice := ui.ShowMainMenue()

        switch choice {

        case 1:
            ui.List_Contacts(contacts)
        case 2:
            contacts = ui.Add_Contacts(contacts)
        case 3:
            err := vault.Save_Vault(vaultPath, contacts, password)
            if err != nil {
                log.Fatalf("Failed to save vault: %v", err)
            }
            fmt.Println("Vault Saved and Locked")
            os.Exit(0)
            
        default:
            fmt.Println("Invalid Option: Try Again")
        }
    }
}