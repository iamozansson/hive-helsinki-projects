package menu

import (
	"fmt"
	"strconv"
)

func OperationSelection() {
	fmt.Println("Select operation (1/2):")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")

	var n int
	for {
		fmt.Print("Your choice: ")
		var input string
		fmt.Scanln(&input)
		var err error
		n, err = strconv.Atoi(input)
		if err == nil && (1 <= n && n <= 2) {
			break
		}
		fmt.Println("Invalid choice.")
	}

	fmt.Printf("You chose option %d\n", n)

	// Call cyper operation menu and passed the opeartion type
	CypherSelection(n)

}

// Return -1 if ivalid operation type
func OperationValidation(operationType string) int {
	operation, error := strconv.Atoi(operationType)
	if operation >= 1 && operation <= 2 && error == nil {
		return operation

	} else {
		PrintOperationInvalidMessage()
		return -1
	}
}

func PrintOperationInvalidMessage() {
	fmt.Println("Ivalid choice for operation type [1/2]")
}
