package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Original Alphabet
var alphabet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// Create 3 Rotors with random Alphabetical Orders
var rotor1 = []rune{'E', 'K', 'M', 'F', 'L', 'G', 'D', 'Q', 'V', 'Z', 'N', 'T', 'O', 'W', 'Y', 'H', 'X', 'U', 'S', 'P', 'A', 'I', 'B', 'R', 'C', 'J'}
var rotor2 = []rune{'D', 'J', 'A', 'K', 'S', 'I', 'R', 'U', 'X', 'B', 'L', 'H', 'W', 'T', 'M', 'C', 'Q', 'G', 'Z', 'N', 'P', 'Y', 'F', 'V', 'O', 'E'}
var rotor3 = []rune{'B', 'D', 'F', 'H', 'J', 'L', 'C', 'P', 'R', 'T', 'X', 'V', 'Z', 'N', 'Y', 'E', 'I', 'W', 'G', 'A', 'K', 'M', 'U', 'S', 'Q', 'O'}

// Create the Reflector with a Static random Alphabet
var reflector = []rune{'T', 'L', 'I', 'H', 'W', 'E', 'K', 'N', 'A', 'Q', 'Z', 'G', 'V', 'P', 'X', 'D', 'O', 'J', 'Y', 'U', 'F', 'S', 'R', 'C', 'M', 'B'}

// Variable to control Printing of Output
var verbose string

// Main Function
func main() {

	// Prompt the User for an Input
	fmt.Println("******************************************************************************************")
	fmt.Println("                                  IST 402 Enigma Machine                                  ")
	fmt.Println("                This Enigma Machine only encrypts English Alphabet Letters                ")
	fmt.Println("                Your String may contain Other Characters, but not numbers!                ")
	fmt.Println("******************************************************************************************")
	fmt.Println("Enter a String to Encrypt:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	// Convert the Input to uppercase
	input = strings.ToUpper(input)

	fmt.Println("-------------------------")
	fmt.Println("Would you like a Detailed Output of the Encryption/Decryption process? (Y/N):")
	fmt.Scanln(&verbose)

	// Print the current Rotor positions
	//fmt.Println(string(rotor1))
	//fmt.Println(string(rotor2))
	//fmt.Println(string(rotor3))

	// Print the Reflector position
	//fmt.Println(string(reflector))

	// Prompt the User for the 3 Encryption Keys
	fmt.Println("-------------------------")
	var encryptionKey1 int
	var encryptionKey2 int
	var encryptionKey3 int
	fmt.Println("Please enter the First Key you wish to use for Encryption (1-26):")
	fmt.Scanln(&encryptionKey1)
	fmt.Println("Please enter the Second Key you wish to use for Encryption (1-26):")
	fmt.Scanln(&encryptionKey2)
	fmt.Println("Please enter the Third Key you wish to use for Encryption (1-26):")
	fmt.Scanln(&encryptionKey3)

	// Encrypt the Input
	encryptedStr := encrypt(input, encryptionKey1, encryptionKey2, encryptionKey3)

	// Prompt the User for the 3 Decryption Keys
	fmt.Println("-------------------------")
	var decryptionKey1 int
	var decryptionKey2 int
	var decryptionKey3 int
	fmt.Println("Please enter the First Key you wish to use for Decryption (1-26):")
	fmt.Scanln(&decryptionKey1)
	fmt.Println("Please enter the Second Key you wish to use for Decryption (1-26):")
	fmt.Scanln(&decryptionKey2)
	fmt.Println("Please enter the Third Key you wish to use for Decryption (1-26):")
	fmt.Scanln(&decryptionKey3)

	// Decrypt the String
	decrypt(encryptedStr, decryptionKey1, decryptionKey2, decryptionKey3)
	fmt.Println("-------------------------")
}

// Encryption Function
func encrypt(str string, r1 int, r2 int, r3 int) string {
	// Check if str length is greater than 0
	if len(str) < 1 {
		return "Invalid Input Length"
	}

	// Check that r1, r2, and r3 are integers between 1 and 26
	if r1 < 1 || r1 > 26 {
		fmt.Println("Error: Rotor 1 Position must be a Positive Integer between 1-26.")
	}
	if r2 < 1 || r2 > 26 {
		fmt.Println("Error: Rotor 2 Position must be a Positive Integer between 1-26.")
	}
	if r3 < 1 || r3 > 26 {
		fmt.Println("Error: Rotor 3 Position must be a Positive Integer between 1-26.")
	}

	// Rotate the Rotors to the specified Positions
	rotate(rotor1, r1)
	//fmt.Println(string(rotor1))
	rotate(rotor2, r2)
	//fmt.Println(string(rotor2))
	rotate(rotor3, r3)
	//fmt.Println(string(rotor3))

	// Remove the Newline Character '\n' and the Return Character '\r' at the end of the String
	output := ""
	for _, char := range str {
		if char != '\n' && char != '\r' {
			output += string(char)
		}
	}
	str = output

	// Split the String Input into a Slice of Characters
	chars := []rune(str)
	for i, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[i] = rune(str[i])
		}
	}

	var greenOutput string
	for _, c := range output {
		greenOutput += fmt.Sprintf("\033[1m\033[32m%c\033[0m", c)
	}

	// Print the Original Mapping
	fmt.Println("-------------------------")
	fmt.Println("String Before Encryption:", string(greenOutput))
	greenOutput = ""

	// Encrypt the String one Character at a time
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 65 && chars[i] <= 90 {
			if verbose == "Y" || verbose == "y" {
				fmt.Println("-------------------------")
			}
			// Map the Letter using Rotor 1
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 1   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor1[chars[i]-'A']))
			}
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
			// Map the Letter using Rotor 2
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 2   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor2[chars[i]-'A']))
			}
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 3
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 3   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor3[chars[i]-'A']))
			}
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using the Reflector
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Reflector - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(reflector[chars[i]-'A']))
			}
			chars[i] = reflector[chars[i]-'A']
			// Map the Letter using Rotor 3
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 3   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor3[chars[i]-'A']))
			}
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using Rotor 2
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 2   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor2[chars[i]-'A']))
			}
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 1
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 1   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor1[chars[i]-'A']))
			}
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
		} else {
			chars[i] = chars[i]
		}
	}

	// Makes the output color green
	output = string(chars)
	for _, c := range output {
		greenOutput += fmt.Sprintf("\033[1m\033[32m%c\033[0m", c)
	}
	// Print the final Mapping
	fmt.Println("-------------------------")
	//fmt.Println("Final Encryption:", string(chars))
	fmt.Println("Final Encryption:", greenOutput)

	// Revert the Rotors to the Original Position after Encrypting
	rotate(rotor1, 27-r1+21)
	rotate(rotor2, 27-r2+21)
	rotate(rotor3, 27-r3+21)

	return string(chars)
}

// Function to Rotate a Rotor by a specified Rotation
func rotate(rotorMap []rune, rotation int) {
	// Reduce Rotation value by 1 because position of 1 would not Rotate the Rotor at all
	rotation = rotation - 1
	rotated := make([]rune, len(rotorMap))

	// Prints the new state of the Rotor
	//fmt.Printf("Now rotating %v by %d.\n", string(rotorMap), rotation)

	// Increment the Position of the Rotor by 1
	for i, c := range rotorMap {
		rotated[(i+rotation)%len(rotorMap)] = c
	}
	copy(rotorMap, rotated)
}

// Decryption Function
func decrypt(str string, r1 int, r2 int, r3 int) string {
	// Check if str length is greater than 0
	if len(str) < 1 {
		return "Invalid Input Length"
	}

	// Check that r1, r2, and r3 are integers between 1 and 26
	if r1 < 1 || r1 > 26 {
		fmt.Println("Error: Rotor 1 Position must be a Positive Integer between 1-26.")
	}
	if r2 < 1 || r2 > 26 {
		fmt.Println("Error: Rotor 2 Position must be a Positive Integer between 1-26.")
	}
	if r3 < 1 || r3 > 26 {
		fmt.Println("Error: Rotor 3 Position must be a Positive Integer between 1-26.")
	}

	// Rotate the Rotors 6 positions to account for the Encryption Positions
	rotate(rotor1, r1+6)
	rotate(rotor2, r2+6)
	rotate(rotor3, r3+6)

	// Remove the Newline Character '\n' and the Return Character '\r' at the end of the String
	output := ""
	for _, char := range str {
		if char != '\n' && char != '\r' {
			output += string(char)
		}
	}
	str = output

	// Split the String Input into a Slice of Characters
	chars := []rune(str)
	for i, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[i] = rune(str[i])
		}
	}

	var greenOutput string
	for _, c := range output {
		greenOutput += fmt.Sprintf("\033[1m\033[32m%c\033[0m", c)
	}

	// Print the Original Mapping
	fmt.Println("-------------------------")
	fmt.Println("String Before Decryption:", greenOutput)
	greenOutput = ""

	// Decrypt the String one Character at a time
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] >= 65 && chars[i] <= 90 {
			if verbose == "Y" || verbose == "y" {
				fmt.Println("-------------------------")
			}
			rotate(rotor1, 26)
			rotorIndex := IndexOf(rotor1, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 1   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotate(rotor2, 26)
			rotorIndex = IndexOf(rotor2, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 2   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotate(rotor3, 26)
			rotorIndex = IndexOf(rotor3, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 3   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotorIndex = IndexOf(reflector, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Reflector - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotate(rotor3, 26)
			rotorIndex = IndexOf(rotor3, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 3   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotate(rotor2, 26)
			rotorIndex = IndexOf(rotor2, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 2   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
			rotate(rotor1, 26)
			rotorIndex = IndexOf(rotor1, chars[i])
			if verbose == "Y" || verbose == "y" {
				fmt.Println("Rotor 1   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			}
			chars[i] = alphabet[rotorIndex]
		} else {
			chars[i] = chars[i]
		}
	}

	// Makes the output color green
	output = string(chars)
	for _, c := range output {
		greenOutput += fmt.Sprintf("\033[1m\033[32m%c\033[0m", c)
	}
	// Print the final Mapping
	fmt.Println("-------------------------")
	//fmt.Println("Final Decryption:", string(chars))

	fmt.Println("Final Decryption:", greenOutput)

	return string(chars)
}

// IndexOf Function to determine the index of a given Slice Element
func IndexOf(haystack []rune, needle rune) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}
