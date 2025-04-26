package vault

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "ModularVault/internal/models"
    "ModularVault/internal/storage"
)


func ListContacts() {
    contacts := storage.LoadContacts()

    if len(contacts) == 0 {
        fmt.Println("No contacts found  :<")
        return
    }
    for _, c := range contacts {
        fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Phone)
    }
}


func AddContact() {

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter name: ")


    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter phone: ")


    phone, _ := reader.ReadString('\n')
    phone = strings.TrimSpace(phone)

    contacts := storage.LoadContacts()
    newID := 1

    if len(contacts) > 0 {
        newID = contacts[len(contacts)-1].ID + 1
    }

    newContact := models.Contact{ID: newID, Name: name, Phone: phone}
    contacts = append(contacts, newContact)

    storage.SaveContacts(contacts)

    fmt.Println("Contact added  :]")
}



func DeleteContact() {
    fmt.Print("Enter ID to delete: ")
    var id int
    fmt.Scanln(&id)

    contacts := storage.LoadContacts()
    updated := []models.Contact{}
    found := false

    for _, c := range contacts {
        if c.ID != id {
            updated = append(updated, c)
        } else {
            found = true
        }
    }

    if !found {
        fmt.Println("Contact not found  :@")
        return
    }

    storage.SaveContacts(updated)
    fmt.Println("Contact deleted.")
}
