package main

import (
	"fmt"
	"os"

	"hive.fi/cypher/menu"
)

func main() {

	withoutArgs := os.Args[1:]

	fmt.Println("Welcome to the Cypher Tool!")

	switch len(withoutArgs) {
	case 0:
		{
			// Open select operation menu
			menu.OperationSelection()
		}
	case 1:
		{

			// Check operation type
			if os.Args[1] != "-h" {
				operationType := menu.OperationValidation(os.Args[1])
				if operationType != -1 {
					menu.CypherSelection(operationType)
				} 
			} else {
				menu.Help()
			}
			
		}
	case 2:
		{
			// Check operation type
			operationType := menu.OperationValidation(os.Args[1])

			// Check Cypher type
			cypherType := menu.CypherValidation(os.Args[2])

			if operationType != -1 && cypherType != -1 {
				menu.EncryptDecrypt(operationType, cypherType)
			}
		}
	default:
		{
			fmt.Println("To many argument !")
		}
	}

}
