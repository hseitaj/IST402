package main

import "fmt"

// Prints a divider bar.
func divider() {
	print("-------------------------\n")
}

// ccEncrypt encrypts a message using the Caesar Cipher and a given shift.
func ccEncrypt(message string, shift int) string {
	cipherText := ""
	for _, char := range message {
		if char >= 'A' && char <= 'Z' {
			// Shift the character by the specified amount
			cipherText += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			cipherText += string(char)
		}
	}
	return cipherText
}

// ccDecrypt decrypts a message using the Caesar Cipher and a given shift.
func ccDecrypt(message string, shift int) string {
	plainText := ""
	for _, char := range message {
		if char >= 'A' && char <= 'Z' {
			// Un-shift the character by the specified amount
			plainText += string((char-'A'-rune(shift)+26)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			plainText += string((char-'a'-rune(shift)+26)%26 + 'a')
		} else {
			plainText += string(char)
		}
	}
	return plainText
}

// ccBruteForce brute forces a Caesar Cipher encrypted message by generating all possible
// shifts and printing the resulting messages.
func ccBruteForce(message string) {
	for i := 0; i < 26; i++ {
		// Call ccDecrypt on the message for each shift value
		bfMessage := ccDecrypt(message, i)
		fmt.Printf("Shift %d: \"%s\"\n", i, bfMessage)
	}
}

func main() {

	// Example usage
	originalMessage := "Hello, World!"
	shiftAmount := 7

	divider()
	fmt.Printf("Original Message: \"%s\"\n", originalMessage)
	divider()

	// Encrypt the message
	encryptedMessage := ccEncrypt(originalMessage, shiftAmount)
	fmt.Printf("Encrypted message: \"%s\"\n", encryptedMessage)
	divider()

	// Brute Force the message
	fmt.Println("Brute Forcing the message...")
	ccBruteForce(encryptedMessage)
	divider()

	// Decrypt the message
	decryptedMessage := ccDecrypt(encryptedMessage, shiftAmount)
	fmt.Printf("Decrypted message: \"%s\"\n", decryptedMessage)
	divider()

}
