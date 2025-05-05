package vault

import (
	"os"
	"errors"
	"encoding/json"
	"path/filepath"

	"WhisperVault/internal/crypto"
	"WhisperVault/internal/models"
)

// Decrypt and Parses the vault file.
func LoadVault(filePath string, password string) ([]models.Contact, error) {
	// Read encrypted vault from disk
	data, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}

	// Decrypt and verify HMAC
	plaintext, err := crypto.Decrypt(data, password)
	if err != nil {
		return nil, err
	}

	// Decode JSON into contacts
	var contacts []models.Contact
	if err := json.Unmarshal(plaintext, &contacts); err != nil {
		return nil, errors.New("vault decrypted, but JSON is invalid")
	}

	return contacts, nil
}

// Encrypt and Write the contact list to disk.
func SaveVault(filePath string, contacts []models.Contact, password string) error {
	// Encode contacts to JSON
	plaintext, err := json.Marshal(contacts)
	if err != nil {
		return err
	}

	// Encrypt with password + HMAC
	ciphertext, err := crypto.Encrypt(plaintext, password)
	if err != nil {
		return err
	}

	// Writing to temp file for safe saving
	tmpPath := filePath + ".tmp"
	if err := os.WriteFile(tmpPath, ciphertext, 0600); err != nil {
		return err
	}

	// Atomic rename to overwrite original vault nicw
	return os.Rename(tmpPath, filePath)
}
