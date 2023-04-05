package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nEnter a message to encrypt: ")
	scanner.Scan()
	msg := scanner.Text()

	fmt.Print("Enter a passphrase (This string is used to generate the encryption key): ")
	scanner.Scan()
	pass := scanner.Text()

	key := sha256.Sum256([]byte(pass))
	aead, _ := chacha20poly1305.NewX(key[:])

	if pass == "" {
		a := make([]byte, 32)
		copy(key[:32], a[:32])
		aead, _ = chacha20poly1305.NewX(a)
	}

	if msg == "" {
		a := make([]byte, 32)
		msg = string(a)
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	ciphertext := aead.Seal(nil, nonce, []byte(msg), nil)

	plaintext, _ := aead.Open(nil, nonce, ciphertext, nil)

	fmt.Println("\n-------------[Starting Encryption...]-------------")
	fmt.Printf("Message: %s\n", msg)
	fmt.Printf("Passphrase: %s\n", pass)
	fmt.Printf("Key: %x\n", key)
	fmt.Printf("Nonce: %x\n", nonce)
	fmt.Printf("Cipher Stream: %x\n", ciphertext)
	fmt.Printf("Decrypted Message: %s\n", plaintext)
	fmt.Println("--------------[Finished Encryption!]--------------")
}
