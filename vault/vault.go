package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io"
	"os"

	"WhisperVault/vault/models"
)

// deriveKey hashes the password using SHA-256 (simple KDF)
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

// Encrypt and save contacts to file
func SaveContacts(filename string, password string, contacts []models.Contact) error {
	// Marshal contacts to JSON
	data, err := json.Marshal(contacts)
	if err != nil {
		return err
	}

	// Generate AES key
	key := deriveKey(password)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use AES-GCM (provides integrity and nonce support)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Create random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Encrypt data
	cipherText := aesGCM.Seal(nonce, nonce, data, nil)

	// Write encrypted data to file
	return os.WriteFile(filename, cipherText, 0644)
}

// Decrypt and load contacts from file
func LoadContacts(filename string, password string) ([]models.Contact, error) {
	var contacts []models.Contact

	// Read file
	cipherText, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Contact{}, nil // No vault yet
		}
		return nil, err
	}

	key := deriveKey(password)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherData := cipherText[:nonceSize], cipherText[nonceSize:]

	// Decrypt
	plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON
	err = json.Unmarshal(plainText, &contacts)
	return contacts, err
}
