package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1  || strings.ToLower(args[0]) == "help" {
		fmt.Println("Usage: ./notestool [TAG]")
		return
	}
	
	name := args[0]
	lines := ReadFile(name)


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\x1b[1mWelcome to the notes tool!\x1b[0m")

	for {
		fmt.Println()
		fmt.Println("\x1b[1mSelect operation:\x1b[0m")
		fmt.Println("\x1b[32m1. \x1b[34mShow notes.\x1b[0m")
		fmt.Println("\x1b[32m2. \x1b[34mAdd a note.\x1b[0m")
		fmt.Println("\x1b[32m3. \x1b[34mDelete a note.\x1b[0m")
		fmt.Println("\x1b[32m4. \x1b[34mExit.\x1b[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" {
			//======================================== List notes ========================================
			fmt.Println()
			lines = ReadFile(name)
			for i, v := range lines {
				fmt.Printf("\x1b[32m%03d\x1b[0m - \x1b[34m%s\x1b[0m\n", i + 1, v)
			}
		} else if input == "2" {
			//======================================== Add note ========================================
			fmt.Println()
			fmt.Println("Enter the note text:")
			
			note, _ := reader.ReadString('\n')
			note = strings.TrimSpace(note)

			if note == "" {
				fmt.Println()
				fmt.Println("\x1b[31mCould not add an empty note\x1b[0m")
				continue
			}
			
			note = fmt.Sprintf("%s - %s", time.Now().Format("15:04:05 02/01/2006"), note)

			fmt.Println()
			fmt.Println("\x1b[32mYour note has been added successfully!\x1b[0m")
			
			lines = append(lines, note)
			UpdateLines(name, lines)
		} else if input == "3" {
			//======================================== Remove note ========================================
			fmt.Println()
			fmt.Println("Enter the number of note to remove or 0 to cancel:")

			option, _ := reader.ReadString('\n')
			option = strings.TrimSpace(option)

			num, err := strconv.Atoi(option)
			if err != nil || num <= 0 || num > len(lines) {
				fmt.Println()
				fmt.Println("\x1b[31mCancelling delete\x1b[0m")
				continue
			}

			fmt.Println()
			fmt.Println("\x1b[31mYour note has been removed!\x1b[0m")

			lines = append(lines[:num-1], lines[num:]...)
			UpdateLines(name, lines)

		} else if input == "4" {
			//======================================== Exit program ========================================
			fmt.Println()
			fmt.Println("\x1b[1mThank you for using Notes Tool app. See you next time!\x1b[0m")
			break
		}
	}
}