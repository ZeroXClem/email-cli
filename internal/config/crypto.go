package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

const keyEnvVar = "EMAIL_CLI_KEY"

func getEncryptionKey() ([]byte, error) {
	key := os.Getenv(keyEnvVar)
	if key == "" {
		return generateKey()
	}
	return base64.StdEncoding.DecodeString(key)
}

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	os.Setenv(keyEnvVar, base64.StdEncoding.EncodeToString(key))
	return key, nil
}

func encrypt(data []byte) (*EncryptedData, error) {
	key, err := getEncryptionKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	cipher := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	return &EncryptedData{
		Data: base64.StdEncoding.EncodeToString(encrypted),
		IV:   base64.StdEncoding.EncodeToString(iv),
	}, nil
}

func decrypt(encData *EncryptedData) ([]byte, error) {
	key, err := getEncryptionKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted, err := base64.StdEncoding.DecodeString(encData.Data)
	if err != nil {
		return nil, err
	}

	iv, err := base64.StdEncoding.DecodeString(encData.IV)
	if err != nil {
		return nil, err
	}

	cipher := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	cipher.XORKeyStream(decrypted, encrypted)

	return decrypted, nil
}