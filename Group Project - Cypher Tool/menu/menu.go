package menu

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"hive.fi/cypher/decryption"
	"hive.fi/cypher/encryption"
)

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func getInput(toEncrypt bool, encoding int, message string) {
	
	result := ""
	// Call Ecryption operation
	if toEncrypt {
		if encoding == 1 {
			result = encryption.Encrypt_rot13(message)

		} else if encoding == 2 {

			result = encryption.Encrypt_reverse(message)
		} else {
			result = encryption.Encrypt_cypher5(message)
		}

		// Call Decryption operation
	} else {
		if encoding == 1 {
			result = decryption.Decrypt_rot13(message)
		} else if encoding == 2 {

			result = decryption.Decrypt_reverse(message)
		} else {
			result = decryption.Decrypt_rot5(message)
		}
	}

	fmt.Printf("\nResult: %s", result)
}


func EncryptDecrypt(operationOption, cypherType int) {

	fmt.Printf("\nYou chose option [%d] for operation\n", operationOption)

	fmt.Printf("You chose option [%d] for Cypher\n", cypherType)

	message := StringPrompt("\nEnter the message:")

	var toEncrypt bool
	if operationOption == 1 {
		toEncrypt = true
	} else {
		toEncrypt = false
	}

	getInput(toEncrypt, cypherType, message)

}

func Help() {
	fmt.Println("NAME")
	fmt.Println("     cypher - Cypher Tool is a simple command-line program written in Go that allows users to encrypt and decrypt messages using different cyphers.")
	fmt.Println("SYNOPSIS")
	fmt.Println("     cypher [OPTION] [OPTION]")
	fmt.Println("DESCRIPTION")
	fmt.Println("     Encrypt and decrypt or vice versa with input message based on selected argument")
	fmt.Println("     1. First argument: Operation type [1/2]")
	fmt.Println("        1 : Encrypt")
	fmt.Println("        2 : Decrypt")
	fmt.Println("     2. Second argument: Cypher type [1/2/3]")
	fmt.Println("        1: ROT13")
	fmt.Println("        2: Reverse")
	fmt.Println("        3: Cypher 5")
	fmt.Println("     -h, display this help and exit ")
	fmt.Println("AUTHOR")
	fmt.Println("     Writen by Nadja Mckenna, Ozan Arikan, Rio Wibowo")
    fmt.Println("GNU coreutils 8.32                                               August 2026\n")
}