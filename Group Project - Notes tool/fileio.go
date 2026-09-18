package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(name string) []string {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}

func UpdateLines(name string, lines []string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < len(lines); i++ {
		_, err = file.WriteString(lines[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		if i != len(lines)-1 {
			_, err = file.WriteString("\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
