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
	_, err = k.Encrypt(testText)
	assert.Nil(t, err)
}

func TestNewKey(t *testing.T) {
	pass := []byte("testPass")
	salt := []byte("testSalt")
	// empty pass
	_, err := mino.NewKey([]byte(""), pass)
	assert.Equal(t, mino.ErrMissKey, err)

	// empty salt
	_, err = mino.NewKey(pass, []byte(""))
	assert.Equal(t, mino.ErrMissSalt, err)

	_, err = mino.NewKey(pass, salt)
	assert.Nil(t, err)
}

func TestDecrypt(t *testing.T) {
	k, err := mino.DefaultKey()
	assert.Nil(t, err)
	testText := []byte("foo")
	errorCipherText := []byte("bar")

	cipherText, err := k.Encrypt(testText)
	assert.Nil(t, err)
	plainText, err := k.Decrypt(cipherText)
	assert.Nil(t, err)
	assert.Equal(t, testText, plainText)

	// error key
	_, err = k.Decrypt(errorCipherText)
	assert.Equal(t, mino.ErrCiphertext, err)

}
