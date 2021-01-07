package mino

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

var (
	// ErrMissKey miss key
	ErrMissKey = errors.New("key is not empty")

	// ErrMissSalt miss salt
	ErrMissSalt = errors.New("salt is not empty")

	// ErrKey invalid key
	ErrKey = errors.New("invalid key")

	// ErrCiphertext invalid ciphertext
	ErrCiphertext = errors.New("invalid ciphertext")
)

const (
	// KeyIter key iter
	KeyIter = 4096

	// KeyLen key length
	KeyLen = 32
)

// Key encrypt and decrypt key
type Key struct {
	Content []byte
}

// NewKey  create new key
func NewKey(passphrase []byte, salt []byte) (key Key, err error) {
	if len(passphrase) < 1 {
		err = ErrMissKey
		return
	}
	if len(salt) < 1 {
		err = ErrMissSalt
		return
	}
	key = Key{
		Content: pbkdf2.Key(passphrase, salt, KeyIter, KeyLen, sha256.New),
	}
	return
}

// DefaultKey returns default key
func DefaultKey() (Key, error) {
	return NewKey([]byte("zsh"), []byte("mini"))
}

// Encrypt encrypt plaintext
func (key Key) Encrypt(plaintext []byte) (ciphertext []byte, err error) {
	ivBytes := make([]byte, 12)
	_, err = rand.Read(ivBytes)
	if err != nil {
		return
	}
	b, err := aes.NewCipher(key.Content)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	ciphertext = aesgcm.Seal(nil, ivBytes, plaintext, nil)
	ciphertext = append(ivBytes, ciphertext...)
	return
}

// Decrypt decrypt ciphertext
func (key Key) Decrypt(ciphertext []byte) (plaintext []byte, err error) {
	if len(ciphertext) < 13 {
		err = ErrCiphertext
		return
	}
	b, err := aes.NewCipher(key.Content)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	plaintext, err = aesgcm.Open(nil, ciphertext[:12], ciphertext[12:], nil)
	return
}
