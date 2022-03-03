package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"
)

func Encrypt(message string) (encodedMsg string, err error) {
	var plainText = []byte(message)

	var block cipher.Block
	var key = []byte(os.Getenv("ENCRYPTION_KEY"))
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	var cipherText = make([]byte, aes.BlockSize+len(plainText))
	var iv []byte = cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(crand.Reader, iv); err != nil {
		log.Panicln(err)
		return
	}

	var stream cipher.Stream = cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encodedMsg = base64.RawStdEncoding.EncodeToString(cipherText)
	return
}

func Decrypt(encodedMsg string) (decodedMsg string, err error) {
	var cipherText []byte
	if cipherText, err = base64.RawStdEncoding.DecodeString(encodedMsg); err != nil {
		return
	}

	var block cipher.Block
	var key = []byte(os.Getenv("ENCRYPTION_KEY"))
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	var iv []byte = cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	var stream cipher.Stream = cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedMsg = string(cipherText)
	return
}