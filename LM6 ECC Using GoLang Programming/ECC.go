package main

// LM6 - ECC in GoLang
// IST 402

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Generate a random curve
	curve := elliptic.P256()

	// Generate a private key
	privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Generate a public key
	//publicKey := elliptic.Marshal(curve, x, y)

	// Get user input
	fmt.Print("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Remove newline character from input
	input = strings.TrimSuffix(input, "\n")

	// Convert input to bytes
	message := []byte(input)

	// Generate a random nonce for AES-GCM encryption
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// Generate shared secret between private and public keys
	sharedSecretX, _ := curve.ScalarMult(x, y, privateKey)
	sharedSecret := sharedSecretX.Bytes()

	// Encrypt the message using AES-GCM
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		panic(err)
	}

	aesGcm, err := cipher.NewGCMWithNonceSize(block, 12)
	if err != nil {
		panic(err)
	}

	ciphertext := aesGcm.Seal(nil, nonce, message, nil)

	// Convert ciphertext to hexadecimal string
	hexCiphertext := hex.EncodeToString(ciphertext)
	fmt.Printf("Encrypted message: %s\n", hexCiphertext)

	// Decrypt the ciphertext using AES-GCM
	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// Print decrypted plaintext
	fmt.Printf("Decrypted message: %s\n", plaintext)
}
