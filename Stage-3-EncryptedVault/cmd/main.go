package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"

    "Stage-3-EncryptedVault/internal/vault"
)

func main() {

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter master password: ")

    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    for {
        fmt.Println("\n-- EncryptedVault --")
        fmt.Println("  1. List Contacts")
        fmt.Println("  2. Add Contact")
        fmt.Println("  3. Delete Contact")
        fmt.Println("  4. Exit")
        fmt.Print("  Choose: ")

        var choice int
        fmt.Scanln(&choice)


        switch choice {

        case 1:
            vault.ListContacts(password)
        case 2:
            vault.AddContact(password)
        case 3:
            vault.DeleteContact(password)
        case 4:
            fmt.Println("  See Ya!  ^-^")
            return

        default:
            fmt.Println("  Invalid option.")
        }
    }
}
