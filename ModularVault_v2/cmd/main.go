package main

import (
    "fmt"
    "ModularVault/internal/vault"
)

func main() {
    for {
        fmt.Println("\n--- ModularVault ---")
        fmt.Println("1. List Contacts")
        fmt.Println("2. Add Contact")
        fmt.Println("3. Delete Contact")
        fmt.Println("4. Exit")
        fmt.Print("Choose option: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            vault.ListContacts()
        case 2:
            vault.AddContact()
        case 3:
            vault.DeleteContact()
        case 4:
            fmt.Println("Thanks for the visit")
            return
        default:
            fmt.Println("Invalid option (^;^)")
        }
    }
}
