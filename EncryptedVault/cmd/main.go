package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "io"
)

//Turning a string password into a 32-byte key using SHA-256
func GenerateKey(password string) []byte {

    hash := sha256.Sum256([]byte(password))
    return hash[:]

}

//Using AES-CTR to encrypt data
func Encrypt(data []byte, password string) ([]byte, error) {

    key := GenerateKey(password)  //Function is above


	//Intialising new AES cipher block with key
	//Return error if AES fails
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }


	//NONCE = Number used ones a unique id we can say 
    nonce := make([]byte, aes.BlockSize)

	//nonce is filled with randomnes and filled with randomness from crypto/rand.Reader
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }


	//using nonce and AES block to create a CTR stream
	//CTR stream - XOR the data with a gen keystream
    stream := cipher.NewCTR(block, nonce)

	//put the slice to hold encypted version of data inputed 
    ciphertext := make([]byte, len(data))

	//XOR's keystream with plaintext to ciphertext
    stream.XORKeyStream(ciphertext, data)


	//prepend the nonce to the ciphertext and return it as [nonce][ciphertext]
    return append(nonce, ciphertext...), nil
}