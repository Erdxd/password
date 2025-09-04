package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"os"
)

type Encrypter struct {
	Key string
}

func NewEncryptor() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("НЕ передан парматер KEY в переменые окружения")
	}
	return &Encrypter{
		Key: key,
	}
}
func (enc *Encrypter) Encrypt(plain []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, plain, nil)
}
func (enc *Encrypter) Decrypt(EncryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := EncryptedStr[:nonceSize], EncryptedStr[:nonceSize]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
