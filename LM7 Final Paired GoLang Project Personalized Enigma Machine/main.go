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

var reflector = []rune{'T', 'L', 'I', 'H', 'W', 'E', 'K', 'N', 'A', 'Q', 'Z', 'G', 'V', 'P', 'X', 'D', 'O', 'J', 'Y', 'U', 'F', 'S', 'R', 'C', 'M', 'B'}

func main() {
	// Prompt the User for an Input
	fmt.Println("Enter a string:")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	// Convert the Input to uppercase
	input = strings.ToUpper(input)

	// Print the current Rotor positions
	//fmt.Println(string(rotor1))
	//fmt.Println(string(rotor2))
	//fmt.Println(string(rotor3))

	// Print the Reflector position
	//fmt.Println(string(reflector))

	// Encrypt the Input
	encryptedStr := encrypt(input, 26, 26, 26)
	decryptedStr := decrypt(encryptedStr, 26, 26, 26)
	print(decryptedStr)
}

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

	// Print the Original Mapping
	fmt.Println("-------------------------")
	fmt.Println("Before Encryption:", string(chars))

	// Encrypt the String one Character at a time
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 65 && chars[i] <= 90 {
			fmt.Println("-------------------------")
			// Map the Letter using Rotor 1
			fmt.Println("Rotor 1   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor1[chars[i]-'A']))
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
			// Map the Letter using Rotor 2
			fmt.Println("Rotor 2   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor2[chars[i]-'A']))
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 3
			fmt.Println("Rotor 3   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor3[chars[i]-'A']))
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using the Reflector
			fmt.Println("Reflector - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(reflector[chars[i]-'A']))
			chars[i] = reflector[chars[i]-'A']
			// Map the Letter using Rotor 3
			fmt.Println("Rotor 3   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor3[chars[i]-'A']))
			chars[i] = rotor3[chars[i]-'A']
			rotate(rotor3, 2)
			// Map the Letter using Rotor 2
			fmt.Println("Rotor 2   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor2[chars[i]-'A']))
			chars[i] = rotor2[chars[i]-'A']
			rotate(rotor2, 2)
			// Map the Letter using Rotor 1
			fmt.Println("Rotor 1   - Remapping Letter:", i+1, "-", string(chars[i]), "->", string(rotor1[chars[i]-'A']))
			chars[i] = rotor1[chars[i]-'A']
			rotate(rotor1, 2)
		} else {
			chars[i] = chars[i]
		}
	}
	// Print the final Mapping
	fmt.Println("-------------------------")
	fmt.Println("Final Encryption:", string(chars))
	fmt.Println("-------------------------")

	for range str {
		rotate(rotor1, 25)
		rotate(rotor2, 25)
		rotate(rotor3, 25)
	}

	return string(chars)
}

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

	rotate(rotor1, r1+(len(str)*2)+1)
	//fmt.Println(string(rotor1))
	rotate(rotor2, r2+(len(str)*2)+1)
	//fmt.Println(string(rotor2))
	rotate(rotor3, r3+(len(str)*2)+1)
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

	// Print the Original Mapping
	fmt.Println("-------------------------")
	fmt.Println("Before Decryption:", string(chars))

	// Decrypt the String one Character at a time
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] >= 65 && chars[i] <= 90 {
			fmt.Println("-------------------------")
			//fmt.Println("Rotor1 Current Position", string(rotor1))
			rotate(rotor1, 26)
			//fmt.Println("Rotor1 New Position    ", string(rotor1))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex := IndexOf(rotor1, chars[i])
			//fmt.Println(string(rotor1[rotorIndex]), "is letter", rotorIndex+1, "in Rotor1")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 1   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Rotor2 Current Position", string(rotor2))
			rotate(rotor2, 26)
			//fmt.Println("Rotor2 New Position    ", string(rotor2))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(rotor2, chars[i])
			//fmt.Println(string(rotor2[rotorIndex]), "is letter", rotorIndex+1, "in Rotor2")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 2   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Rotor3 Current Position", string(rotor3))
			rotate(rotor3, 26)
			//fmt.Println("Rotor3 New Position    ", string(rotor3))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(rotor3, chars[i])
			//fmt.Println(string(rotor3[rotorIndex]), "is letter", rotorIndex+1, "in Rotor3")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 3   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Reflector Current Position", string(reflector))
			//fmt.Println("Reflector New Position    ", string(reflector))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(reflector, chars[i])
			//fmt.Println(string(reflector[rotorIndex]), "is letter", rotorIndex+1, "in Reflector")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Reflector - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Rotor3 Current Position", string(rotor3))
			rotate(rotor3, 26)
			//fmt.Println("Rotor3 New Position    ", string(rotor3))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(rotor3, chars[i])
			//fmt.Println(string(rotor3[rotorIndex]), "is letter", rotorIndex+1, "in Rotor3")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 3   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Rotor2 Current Position", string(rotor2))
			rotate(rotor2, 26)
			//fmt.Println("Rotor2 New Position    ", string(rotor2))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(rotor2, chars[i])
			//fmt.Println(string(rotor2[rotorIndex]), "is letter", rotorIndex+1, "in Rotor2")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 2   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
			//fmt.Println("Rotor1 Current Position", string(rotor1))
			rotate(rotor1, 26)
			//fmt.Println("Rotor1 New Position    ", string(rotor1))
			//fmt.Println("Chars[i]:", chars[i], "=", string(chars[i]))
			rotorIndex = IndexOf(rotor1, chars[i])
			//fmt.Println(string(rotor1[rotorIndex]), "is letter", rotorIndex+1, "in Rotor1")
			//fmt.Println("So that means", string(chars[i]), "was mapped to", string(alphabet[rotorIndex]))
			fmt.Println("Rotor 1   - Demapping Letter:", i+1, "-", string(chars[i]), "->", string(alphabet[rotorIndex]))
			chars[i] = alphabet[rotorIndex]
			//fmt.Println("-------------------------")
		} else {
			chars[i] = chars[i]
		}
	}

	// Print the final Mapping
	fmt.Println("-------------------------")
	fmt.Println("Final Decryption:", string(chars))
	fmt.Println("-------------------------")

	return string(chars)
}

func IndexOf(haystack []rune, needle rune) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}
