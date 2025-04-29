package vault

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"WhisperVault_v5/internal/models"
	"WhisperVault_v5/internal/storage"
)

func ListContacts(password string) {

	contacts, err := storage.LoadContacts(password)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Name, c.Phone)
	}
}

func AddContact(password string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	contacts, _ := storage.LoadContacts(password)

	newID := 1

	if len(contacts) > 0 {
		newID = contacts[len(contacts)-1].ID + 1
	}

	newContact := models.Contact{ID: newID, Name: name, Phone: phone}
	contacts = append(contacts, newContact)

	err := storage.SaveContacts(contacts, password)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Contact Saved Succesfully")
	}
}

func DeleteContact(password string) {

	var id int
	fmt.Print("Enter ID to delete: ")
	fmt.Scanln(&id)

	contacts, err := storage.LoadContacts(password)

	if err != nil {
		fmt.Println("  Error: ", err)
		return
	}

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
		fmt.Println("  Contact not found")
		return
	}

	err = storage.SaveContacts(updated, password)

	if err != nil {
		fmt.Println("  Error: ", err)
	} else {
		fmt.Println("  Deletion Succussful")
	}

}
