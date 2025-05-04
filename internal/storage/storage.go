package storage

import (
	"os"
	"errors"
	"enooding/json"
	
	"WhisperVault/internal/crypto"
	"WhisperVault/internal/models"
)

// SaveContacts encrypts the given contact list using the provided password,
// then writes the resulting vault blob to disk.
func SaveContacts(contacts []models.Contact, password string) error {
	
	// Marshal contacts to JSON
	data, err := json.Marshal(contacts)
	if err != nil {
		return err
	}

	// Encrypt the data using our crypto layer (PBKDF2 + AES-GCM + HMAC)
	encryptedBlob, err := crypto.Encrypt(data, password)
	if err != nil {
		return err
	}

	// Write encrypted blob to file
	err = os.WriteFile("encrypted_contacts.vault", encryptedBlob, 0600)
	if err != nil {
		return err
	}

	return nil
}

// LoadContacts reads the vault file, verifies HMAC and decrypts using password.
// Returns the contact list if successful, or an error if password is wrong or file is corrupted.
func LoadContacts(password string) ([]models.Contact, error) {
	// Read the vault file from disk
	blob, err := os.ReadFile("encrypted_contacts.vault")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If the vault doesn't exist yet, return an empty contact list
			return []models.Contact{}, nil
		}
		return nil, err
	}

	// Decrypt and verify HMAC using the given password
	plaintext, err := crypto.Decrypt(blob, password)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON back into contact list
	var contacts []models.Contact
	err = json.Unmarshal(plaintext, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}
