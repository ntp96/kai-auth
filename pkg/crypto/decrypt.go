package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func Decrypt(cipheredText []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipheredText) < nonceSize {
		return nil, errors.New("cipheredText too short")
	}

	nonce, cipheredText := cipheredText[:nonceSize], cipheredText[nonceSize:]
	return gcm.Open(nil, nonce, cipheredText, nil)
}