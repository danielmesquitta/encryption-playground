package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type AES struct {
	key []byte
}

// NewAES creates a new AES instance with a 32-byte key.
func NewAES() *AES {
	key := []byte(os.Getenv("AES_KEY"))
	if key == nil {
		panic("AES_KEY environment variable is not set")
	}
	if len(key) != 32 {
		panic("AES_KEY must be 32 bytes")
	}
	return &AES{key: key}
}

// Encrypt locks up your plaintext with a 32-byte key.
func (a *AES) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt reverses what Encrypt did.
func (a *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ct := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ct, nil)
}
