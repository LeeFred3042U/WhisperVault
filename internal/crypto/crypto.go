package crypto

import(
	"io"
	"errors"
	"bytes"
	"encoding/binary"

	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)


const (
	saltSize       = 	16            
	keySize        = 	32            
	nonceSize      = 	12            
	hmacSize       = 	32            
	iterations     = 	600_000       
	magicHeader    = 	"WVAULT"      
	versionByte    = 	byte(1)       
)


// Generate a Key
func GenerateKey(password string, salt []byte) []byte {
	return pdkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

func Encrypt(plaintext []byte, password string) ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil{ return nil, err}
	key := DeriveKey(password, salt)

	// AES-GCM
	block, err := aes.NewCipher(key); if err != nil {return nil, err}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, nonceSize); if err != nil { return nil, err}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// FINAL blob: [magic][version][salt][nonce][ciphertext][hmac]
	buf := bytes.Buffer{}
	buf.Write([]byte(magicHeader))
	buf.WriteByte(versionByte)
	buf.Write(salt)
	buf.Write(nonce)
	buf.Write(ciphertext)

	// Compute HMAC
	mac := hmac.New(sha256.New, key)
	mac.Write(buf.Bytes())
	hmacSum := mac.Sum(nil)

	// Final Blob = data + HMAC
	buf.Write(hmacSum)
	return buf.Bytes(), nil
}