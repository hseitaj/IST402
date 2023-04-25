package main

// LM6 - ECC in GoLang
// IST 402

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
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
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Generate a public key
	publicKey := &privateKey.PublicKey

	// Get user input
	print("-------------------------\n")
	fmt.Print("Enter a string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	print("-------------------------\n")
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
	sharedSecretX, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
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
	print("-------------------------\n")
	// Decrypt the ciphertext using AES-GCM
	plaintext, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	// Print decrypted plaintext
	fmt.Printf("Decrypted message: %s\n", plaintext)
	print("-------------------------\n")
}
