package vault

import (
	"crypto"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"encoding.json"

	"WhisperVault/internal/crypto"
	"Whispervault/internal/models"
)

// Loading and decrypting Vault File
func LoadVault(filePath string, password string) ([]models.Contact, error) {

	// Reading Vault
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	//Decrpt.Verify_HMAC
	plaintext, err := crypto.Decrypt(data, password)
	if err != nil {
		return nil, err
	}

	// Decoding jSON
	var contacts []models.Contact
	err = json.Unmarshal(plaintext, &contacts)
	if err != nil {
		return nil, errors.New("vault derypted but JSON invalid")
	}
	return contacts, nil
}

// Saving and encrypting contact data
func Save_Vault(filePath string, contacts []models.Contact, password string) error {

	// Encoding Contacts
	plaintext, err := json.Marshal(contacts)
	if err != nil {
		return err
	}

	// Encryt + HMAC
	ciphertext, err := crypto.Encrypt(plaintext, password)
	if err != nil {
		return err
	}

	// atomic saving
	tmpPath := filePath + ".tmp"
	err = os.WriteFile(tmpPath, ciphertext, 0600)
	if err != nil {
		return err
	}

	// Atomic Swap
	return os.Rename(tmpPath, filePath)
}