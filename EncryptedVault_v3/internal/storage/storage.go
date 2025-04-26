package storage

import (
	"os"
	"fmt"
	"encoding/json"

	"EncryptedVault_v3/internal/crypto"
	"EncryptedVault_v3/internal/models"
)

const VaultFile = "encrypted_contacts.vault"

func LoadContacts(password string) ([]models.Contact, error) {
    data, err := os.ReadFile(VaultFile)
    if err != nil {
        fmt.Println("  LOAD ERROR:", err)
        return nil, err
    }

    decrypted, err := crypto.Decrypt(data, password)
    if err != nil {
        fmt.Println("  DECRYPT ERROR:", err)
        return nil, err
    }

    var contacts []models.Contact
    err = json.Unmarshal(decrypted, &contacts)
    if err != nil {
        fmt.Println("  UNMARSHAL ERROR:", err)
    }
    return contacts, err
}



func SaveContacts(contacts []models.Contact, password string) error {
	jsonData, err := json.MarshalIndent(contacts, "", "  ")

	if err != nil {
		return err
	}

	encrypted, err := crypto.Encrypt(jsonData, password)

	if err != nil {
		return err
	}


	return os.WriteFile(VaultFile, encrypted, 0644)
}