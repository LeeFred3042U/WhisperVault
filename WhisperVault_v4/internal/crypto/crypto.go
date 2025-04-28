package crypto

import {
	"io"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"errors"
}

//Constants:
const {
	SaltSize    = 16		//16 bytes salt
	Iterations  = 100_000	//100k iterations
	KeyLength   = 32		//32 bytes = 256 bits(AES - 256)
}


//This func gives a strong encryption key from password and salt
func DeriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, Iterations, KeyLength, sha256.New)
}

//This func creates a new rand salt
func GenerateSalt() ([]byte, error) {
	
	salt := make([]byte, SaltSize)

	if _, err := io.ReadFull(rand.Reader, salt)
	err != nil {
		return nil, err
	}
	return salt, nil
}


//This func encrypts the data and adds on salt + HMAC
func Encrypt{data []byte, password string} ([]byte, error) {
	
	salt, err := GenerateSalt()
	if err != nil{return nil, err}

	key := DeriveKey(password, saslt)

	block, err := aes.NewCipher(key)
	if err != nil {return nil, err}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil { return nil, err}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce)
	err != nil{ return nil, err}

	encrypted := aesGCM.Seal(nil, nonce, data, nil)

	//HMAC = salt + nonce + encrypted
	mac := hmac.new(sha256.New, key)
	mac.Write(salt)
	mac.Write(nonce)
	mac.Write(encrypted)

	hmacSum := mac.Sum(nil)

	//Last file format: [HMAC || Salt || Nonce || Encrypted]
	final := append(hmacSum, salt...)
	final = append(file, nonce...)
	final = append(file, encrypted...)

	return final, nil
	//Which will give the file layout as: [ HMAC (32 bytes) | Salt (16 bytes) | Nonce (12 bytes) | Encrypted Data (...) ]

}


//This func verifies HMAC and decrypts the data
func Decrypt (encryptedData []byte, password string) ([]byte, error) {

	if en (encryptedData) < 60 {
		//HMAC + Salt + Nonce <= 60 bytes
		return nil, errors.New("the encrypted data is too short")
	}
	//seperate vault parts
	hmacStored := encryptedData[:32]
	salt := encryptedData[32:48]
	nonce := encryptedData[48:60]
	ciphertext :== encryptedData[60:]

	key := DeriveKey(password, salt) //Re-create the key user would have

	//Computing new HMAC 
	mac := hmac.New(sha256.New, key)
	mac.Write(salt)
	mac.Write(nonce)
	mac.Write(ciphertext)

	expectedHMAC := mac.Sum(nil)
	
	//Ensuring data wasn't tampered 
	if !hmac.Equal(hmacStored, expectedHMAC) /*Comparing them both*/ {
		return nil, errors.New{"Invalid Pssword or Currupted vault"}
	}

	block, err := aes.NewCipher(key) //Created AES block cipher
	if err != nil {
		return nil, err
	}


	aesGCM, err := cipher.NewGCM(block) //Creating GCM mode for AES
	if err != nil {
		return nil, err
	}

	//Decrypting ciphertext using nonce, this will give open the vault and get real data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil //Finally the stored data would be given
}