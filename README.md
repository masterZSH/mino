<image width="64" height="64" src="./images/icon.png" />


## Instruction

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

## Examples

<a href="./examples/chat/README.md">1v1 command-line chat</a>


## Doc
[doc](https://godoc.org/github.com/masterZSH/mino)


## Ctl

[minoctl](https://github.com/masterZSH/mino/tree/main/cli) 
