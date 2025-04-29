package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "WhisperVault_v5/internal/storage"
    "WhisperVault_v5/internal/models"
    "golang.org/x/term"
    "syscall"
)

func main() {
    var contacts []models.Contact
    var password string

    maxRetries := 3
    for attempts := 0; attempts < maxRetries; attempts++ {
        pw, _ := AskPassword("Enter vault password: ")

        cts, err := storage.LoadContacts(pw)
        if err == nil {
            fmt.Println("\nVault unlocked successfully!")
            contacts = cts
            password = pw
            break
        } else {
            fmt.Println("\nInvalid password. Try again.")
        }

        if attempts == maxRetries-1 {
            fmt.Println("\nToo many failed attempts. Exiting.")
            os.Exit(1)
        }
    }

    for {
        fmt.Println("\n--- WhisperVault Menu ---")
        fmt.Println("1. List Contacts")
        fmt.Println("2. Add New Contact")
        fmt.Println("3. Save and Exit")

        fmt.Print("Enter choice: ")
        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            listContacts(contacts)
        case 2:
            contacts = addContact(contacts)
            storage.SaveContacts(contacts, password)
            fmt.Println("\nVault saved after adding contact.")
        case 3:
            storage.SaveContacts(contacts, password)
            fmt.Println("\nVault saved. Goodbye!")
            os.Exit(0)
        default:
            fmt.Println("\nInvalid option. Try again.")
        }
    }
}

// Read password hidden from screen
func AskPassword(prompt string) (string, error) {
    fmt.Print(prompt)
    passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
    fmt.Println()
    if err != nil {
        return "", err
    }
    return string(passwordBytes), nil
}

// List all contacts
func listContacts(contacts []models.Contact) {
    if len(contacts) == 0 {
        fmt.Println("\nNo contacts saved.")
        return
    }
    for i, c := range contacts {
        fmt.Printf("\n%d. Name: %s\n", i+1, c.Name)
        fmt.Printf("   Phone: %s\n", c.Phone)
        fmt.Printf("   Email: %s\n", c.Email)
    }
}

// Add a new contact
func addContact(contacts []models.Contact) []models.Contact {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("\nEnter Name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter Phone: ")
    phone, _ := reader.ReadString('\n')
    phone = strings.TrimSpace(phone)

    fmt.Print("Enter Email: ")
    email, _ := reader.ReadString('\n')
    email = strings.TrimSpace(email)

    newContact := models.Contact{
        Name:  name,
        Phone: phone,
        Email: email,
    }

    contacts = append(contacts, newContact)
    fmt.Println("\nContact added successfully!")
    return contacts
}
