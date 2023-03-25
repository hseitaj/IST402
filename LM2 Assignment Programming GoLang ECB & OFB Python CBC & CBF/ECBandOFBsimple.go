package main

import (
	"fmt"
)

var codebook = [4][2]int{{0b000, 0b100}, {0b010, 0b101}, {0b101, 0b111}, {0b111, 0b001}}

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

// In the ECB implementation, we define a function which takes an array of plaintext blocks
// and returns an array of corresponding ciphertext blocks. We then iterate through each block of the plaintext,
// look up the corresponding ciphered value using the codebookLookup() function, and append it to the ciphertext array.
// Finally, we output the resulting ciphertext in binary format.
func encryptECB(plaintext []int) []int {
	var ciphertext []int
	for _, block := range plaintext {
		lookupValue := codebookLookup(block)
		ciphertext = append(ciphertext, lookupValue)
	}
	return ciphertext
}

// This implementation uses the initialization vector as the feedback value for the first block,
// and then updates the feedback value based on the previous block's feedback and the current plaintext block.
func encryptOFB(plaintext []int, iv int) []int {
	var ciphertext []int
	var feedback int = iv
	for _, block := range plaintext {
		lookupValue := codebookLookup(feedback)
		ciphertext = append(ciphertext, lookupValue)
		feedback = codebookLookup(feedback)
		feedback ^= block
	}
	return ciphertext
}

func main() {
	// Example plaintext
	plaintext := []int{0b000, 0b010, 0b101, 0b111}
	fmt.Printf("The plaintext: %b\n", plaintext)

	// Encrypt using ECB mode
	ciphertext := encryptECB(plaintext)

	// Output ciphered values
	fmt.Printf("ECB ciphertext: %b\n", ciphertext)

	// Example initialization vector
	iv := 0b10

	// Encrypt using OFB mode
	ciphertextOFB := encryptOFB(plaintext, iv)

	// Output ciphered values
	fmt.Printf("OFB ciphertext: %b\n", ciphertextOFB)

}
