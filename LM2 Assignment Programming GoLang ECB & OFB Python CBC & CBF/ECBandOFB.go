package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

// ECB - This program uses the Go standard library's crypto/aes package to create an AES cipher block
// and then encrypts the plaintext using ECB mode.
func main() {
	// Generate a random 128-bit key
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Create the AES cipher block
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Accept plaintext input from the user
	fmt.Print("Enter plaintext: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plaintext := scanner.Text()

	// Pad the plaintext to a multiple of the block size
	blockSize := cipherBlock.BlockSize()
	plaintext = pad(plaintext, blockSize)

	// Create the ECB mode encrypter
	encrypter := NewECBEncrypter(cipherBlock)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	encrypter.CryptBlocks(ciphertext, []byte(plaintext))

	// Print the plaintext, key, and ciphertext details
	fmt.Printf("\nECB implementation:\n")
	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("AES-128 key: %x\n", key)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Create a random IV
	iv := make([]byte, cipherBlock.BlockSize())
	_, err = rand.Read(iv)
	if err != nil {
		panic(err)
	}

	// Create the OFB mode encrypter
	encrypterOFB := cipher.NewOFB(cipherBlock, iv)

	// Encrypt the plaintext
	ciphertextOFB := make([]byte, len(plaintext))
	encrypterOFB.XORKeyStream(ciphertextOFB, []byte(plaintext))

	// Print the plaintext, key, IV, and ciphertext details
	fmt.Printf("\nOFB implementation:\n")
	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("AES-128 key: %x\n", key)
	fmt.Printf("IV: %x\n", iv)
	fmt.Printf("Ciphertext: %x\n", ciphertextOFB)

}

// pad pads the input to a multiple of the block size using PKCS#7 padding
func pad(input string, blockSize int) string {
	padding := blockSize - len(input)%blockSize
	padtext := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padtext[i] = byte(padding)
	}
	return input + string(padtext)
}

// NewECBEncrypter creates a new ECB mode encrypter
func NewECBEncrypter(block cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{block}
}

type ecbEncrypter struct {
	b cipher.Block
}

func (x *ecbEncrypter) BlockSize() int { return x.b.BlockSize() }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.b.BlockSize() != 0 {
		panic("input not full blocks")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.b.BlockSize()])
		src = src[x.b.BlockSize():]
		dst = dst[x.b.BlockSize():]
	}
}
