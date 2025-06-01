package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2 struct {
	pepper []byte
}

// NewArgon2 creates a new Argon2 instance.
func NewArgon2() *Argon2 {
	pepper := []byte(os.Getenv("PASSWORD_PEPPER"))
	if pepper == nil {
		panic("PASSWORD_PEPPER environment variable is not set")
	}

	return &Argon2{
		pepper: pepper,
	}
}

// HashPasswordArgon2id returns a base64-encoded salt+hash blob.
func (a *Argon2) HashPassword(plain string) (string, error) {
	// Argon2 parameters—tune to your environment
	const (
		time    = 1
		memory  = 64 * 1024
		threads = 4
		keyLen  = 32
		saltLen = 16
	)

	// 1. Generate a random salt
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("salt generation failed: %w", err)
	}

	// 2. Append pepper to password
	p := append([]byte(plain), a.pepper...)

	// 3. Derive key
	key := argon2.IDKey(p, salt, time, memory, threads, keyLen)

	// 4. Encode salt||key for storage
	b64 := base64.RawStdEncoding
	encoded := fmt.Sprintf("%s.%s", b64.EncodeToString(salt), b64.EncodeToString(key))
	return encoded, nil
}

// ComparePasswordArgon2id verifies a plaintext against a stored salt+hash.
func (a *Argon2) ComparePassword(encoded, plain string) bool {
	parts := strings.Split(encoded, ".")
	if len(parts) != 2 {
		return false
	}
	b64 := base64.RawStdEncoding
	salt, err1 := b64.DecodeString(parts[0])
	hash, err2 := b64.DecodeString(parts[1])
	if err1 != nil || err2 != nil {
		return false
	}

	// Re-derive key with same params
	p := append([]byte(plain), a.pepper...)
	derived := argon2.IDKey(p, salt, 1, 64*1024, 4, uint32(len(hash)))

	// Constant-time compare
	return subtle.ConstantTimeCompare(hash, derived) == 1
}
