package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize    = 16               // 128-bit random salt
	keySize     = 32               // AES-256 = 32 bytes key
	nonceSize   = 12               // AES-GCM standard nonce size
	hmacSize    = 32               // SHA-256 HMAC output size
	iterations  = 600_000          // PBKDF2 strength (stretching)
	magicHeader = "WVAULT"         // Vault file magic header
	versionByte = byte(1)          // Format version = 1
)

// This applies PBKDF2 to password+salt for secure key
func DeriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

// Encrypt performs AES-GCM encryption + appends HMAC for integrity
func Encrypt(plaintext []byte, password string) ([]byte, error) {
	// Generate a new salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	key := DeriveKey(password, salt)

	// Create AEAD cipher (AES-GCM)
	aesgcm, err := initAEAD(key)
	if err != nil {
		return nil, err
	}

	// Generate a new nonce
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt plaintext
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// Assemble the vault file: [magic][version][salt][nonce][ciphertext]
	buf := bytes.Buffer{}
	buf.Write([]byte(magicHeader)) // eg:- WVAULT
	buf.WriteByte(versionByte)     // Format version
	buf.Write(salt)                // Random salt
	buf.Write(nonce)               // Random nonce
	buf.Write(ciphertext)          // Encrypted content

	// Compute HMAC over everything so far
	mac := hmac.New(sha256.New, key)
	mac.Write(buf.Bytes())
	hmacSum := mac.Sum(nil)

	// Append HMAC at the end
	buf.Write(hmacSum)

	return buf.Bytes(), nil
}

// initAEAD initializes AES-GCM with given key
func initAEAD(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCMWithNonceSize(block, nonceSize)
}

// Decrypt verifies HMAC and decrypts the vault content
func Decrypt(blob []byte, password string) ([]byte, error) {
	// Sanity check: is vault file large enough to contain all parts?
	minSize := len(magicHeader) + 1 + saltSize + nonceSize + hmacSize
	if len(blob) < minSize {
		return nil, errors.New("vault too small or corrupt")
	}


	//Now Verifing 



	// Verify magic header
	if string(blob[:len(magicHeader)]) != magicHeader {
		return nil, errors.New("invalid vault file format")
	}

	// Verify version
	if blob[len(magicHeader)] != versionByte {
		return nil, errors.New("unsupported vault version")
	}

	// Split parts
	offset := len(magicHeader) + 1
	salt := blob[offset : offset+saltSize]
	offset += saltSize

	nonce := blob[offset : offset+nonceSize]
	offset += nonceSize

	hmacStart := len(blob) - hmacSize
	ciphertext := blob[offset:hmacStart]
	expectedHMAC := blob[hmacStart:]

	// Derive key again from password + salt
	key := DeriveKey(password, salt)

	// Verify HMAC
	mac := hmac.New(sha256.New, key)
	mac.Write(blob[:hmacStart])
	if !hmac.Equal(mac.Sum(nil), expectedHMAC) {
		return nil, errors.New("HMAC verification failed – wrong password or tampered vault")
	}

	// Now Decrypt
	aesgcm, err := initAEAD(key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize AEAD: %w", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed")
	}

	return plaintext, nil
}
