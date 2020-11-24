package mino_test

import (
	"testing"

	"github.com/masterZSH/mino"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	k, err := mino.DefaultKey()
	assert.Nil(t, err)
	testText := []byte("foo")
	cipherText, err := k.Encrypt(testText)
	assert.Nil(t, err)
	plainText, err := k.Decrypt(cipherText)
	assert.Nil(t, err)
	assert.Equal(t, testText, plainText)
}

func TestNewKey(t *testing.T) {
	pass := []byte("testPass")
	salt := []byte("testSalt")
	_, err := mino.NewKey(pass, salt)
	assert.Nil(t, err)
}
