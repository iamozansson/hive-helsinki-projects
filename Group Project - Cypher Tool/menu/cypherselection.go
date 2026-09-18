package menu

import (
	"fmt"
	"strconv"

	"hive.fi/cypher/decryption"
	"hive.fi/cypher/encryption"
)

func CypherSelection(operationOption int) {

	fmt.Println("\nSelect cypher (1/2/3)")
	fmt.Println("1. ROT13")
	fmt.Println("2. Reverse")
	fmt.Println("3. Cypher 5")

	var n int
	for {
		fmt.Print("Your choice: ")
		var input string
		fmt.Scanln(&input)
		var err error
		n, err = strconv.Atoi(input)
		if err == nil && (1 <= n && n <= 3) {
			break
		}
		fmt.Println("Invalid choice.")
	}
	fmt.Printf("You chose option %d\n", n)

	message := StringPrompt("\nEnter the message:")

	result := ""
	// Call Ecryption operation
	if operationOption == 1 {
		if n == 1 {
			result = encryption.Encrypt_rot13(message)

		} else if n == 2 {

			result = encryption.Encrypt_reverse(message)
		} else {
			result = encryption.Encrypt_cypher5(message)
		}

		// Call Decryption operation
	} else if operationOption == 2 {
		if n == 1 {
			result = decryption.Decrypt_rot13(message)
		} else if n == 2 {

			result = decryption.Decrypt_reverse(message)
		} else {
			result = decryption.Decrypt_rot5(message)
		}
	}

	fmt.Println("\nResult: ", result)

}

// Return -1 if ivalid Cypher type
func CypherValidation(cypherType string) int {
	cypher, error := strconv.Atoi(cypherType)
	if cypher >= 1 && cypher <= 3 && error == nil {
		return cypher

	} else {
		PrintCypherInvalidMessage()
		return -1
	}
}

func PrintCypherInvalidMessage() {
	fmt.Println("Ivalid choice for Cypher type [1/2/3]")
}
