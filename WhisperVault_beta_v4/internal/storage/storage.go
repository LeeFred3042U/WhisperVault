package storage

import (
	"os"
	"errors"
	"encoding/json"
	"WhisperVault_beta_v4/internal/crypto"
	"WhisperVault_beta_v4/internal/models"
)

const VaultFile = "encrypted_contacts.vault"

func SaveContacts(contacts []models.Contact, password string) error {

	data, err := json.Marshal(contacts)
	if err != nil {
		return nil
	}

	encrypted, err := crypto.Encrypt(data, password)
	if err != nil {
		return err
	}

	return os.WriteFile(VaultFile, encrypted, 0600)
}


func LoadContacts(password string) ([]models.Contact, error) {

	encrypted, err := os.ReadFile(VaultFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Contact{}, nil
		}
		return nil, err
	}

	data, err := crypto.Decrypt(encrypted, password) 
	if err != nil {
		return nil, err
	}

	var contacts []models.Contact
    if err := json.Unmarshal(data, &contacts)
	 err != nil {
        return nil, errors.New("vault decryption succeeded but data corrupted")
    }
	return contacts, nil
}