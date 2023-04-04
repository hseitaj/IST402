package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// read the user's input from the command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your message: ")
	input, _ := reader.ReadString('\n')

	// remove the newline character from the user's input
	input = input[:len(input)-1]

	// perform TLS encryption on the user's input
	fmt.Println("Encrypting message...")
	encrypted := make([]byte, len(input))
	copy(encrypted, input)
	conn.Write(encrypted)

	// perform TLS decryption on the client
	fmt.Println("Decrypting message...")
	decrypted := make([]byte, len(input))
	conn.Write(decrypted)

	// display the decrypted message
	fmt.Println("Decrypted message:", string(decrypted))
}
