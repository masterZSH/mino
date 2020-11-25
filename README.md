# Instruction

Data encryption and Data decryption.

It can already be safely transferred across all systems.

## Install

```
go get -u github.com/masterZSH/mino
```

## Usage

```go

    // create a new key by pass and salt
	k, err := mino.NewKey("myPass", "mySalt")

    // Encrypt plaintext
    cipherText, err := k.Encrypt(plainText)

    // Decrypt ciphertext
    plainText, err := k.Decrypt(cipherText)

```
