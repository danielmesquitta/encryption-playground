package bcrypt

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	pepper []byte
}

func NewBcrypt() *Bcrypt {
	pepper := []byte(os.Getenv("PASSWORD_PEPPER"))
	if pepper == nil {
		panic("PASSWORD_PEPPER environment variable is not set")
	}

	return &Bcrypt{
		pepper: pepper,
	}
}

// HashPassword works with bcrypt for secure password storage.
func (b *Bcrypt) HashPassword(
	plain string,
) (string, error) {
	p := append([]byte(plain), b.pepper...)
	hash, err := bcrypt.
		GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt hash failed: %w", err)
	}
	return string(hash), nil
}

// CheckPassword compares a plaintext password with its bcrypt hash.
func (b *Bcrypt) ComparePassword(
	storedBcrypt, plain string,
) bool {
	p := append([]byte(plain), b.pepper...)
	err := bcrypt.
		CompareHashAndPassword([]byte(storedBcrypt), p)
	return err == nil
}
