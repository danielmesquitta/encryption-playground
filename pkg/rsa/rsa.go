package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type RSA struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewRSA() (*RSA, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return &RSA{PrivateKey: priv, PublicKey: &priv.PublicKey}, nil
}

func (r *RSA) Encrypt(msg []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, r.PublicKey, msg, nil)
}

func (r *RSA) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, r.PrivateKey, ciphertext, nil)
}
