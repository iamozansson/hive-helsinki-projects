package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Weather struct {
	id    int
	key   string
	value float64
	check bool
}

func main() {
	wData := []Weather{
		{id: 1, key: "airTemp", value: 0, check: false},
		{id: 2, key: "airPressure", value: 0, check: false},
		{id: 7, key: "precipitation", value: 0, check: false},
		{id: 11, key: "windSpeed", value: 0, check: false},
		{id: 12, key: "windDirection", value: 0, check: false},
		{id: 13, key: "humidity", value: 0, check: false},
		{id: 14, key: "dewPoint", value: 0, check: false},
		{id: 15, key: "soilMoisture", value: 0, check: false},
		{id: 22, key: "cloudCover", value: 0, check: false},
	}

	fmt.Println("--- Weather Station ---")

	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		choose := userInput
		var id int
		var value float64
		var strVal string

		if strings.Contains(userInput, ",") {
			arr := strings.Split(userInput, ",")

			if len(arr) != 2 {
				fmt.Println("Error: Invalid format. Expected id, value")
				continue
			}

			var firstErr error
			id, firstErr = strconv.Atoi(strings.TrimSpace(arr[0]))

			if firstErr != nil {
				fmt.Printf("Error: ID '%s' is not a valid number\n", arr[0])
				continue
			}

			strVal = strings.TrimSpace(arr[1])

			var secErr error

			if strVal != "NULL" {
				value, secErr = strconv.ParseFloat(strVal, 64)
			}

			if secErr != nil {
				fmt.Printf("Error: Value '%s' is not a valid number\n", strVal)
				continue
			}

			choose = "set"
		}

		switch choose {
		case "set":
			if strVal == "NULL" {
				wData = SetNull(wData, id)
			} else {
				wData = Set(wData, id, value)
			}

		case "get":
			Get(wData)

		case "clear":
			wData = Clear(wData)

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown command.")
		}
	}
}

func Set(wData []Weather, id int, value float64) []Weather {
	for i := 0; i < len(wData); i++ {
		if wData[i].id == id {
			wData[i].value = value
			wData[i].check = true
		}
	}
	return wData
}

func SetNull(wData []Weather, id int) []Weather {
	for i := 0; i < len(wData); i++ {
		if wData[i].id == id {
			wData[i].check = false
		}
	}
	return wData
}

func Get(wData []Weather) {
	for i := 0; i < len(wData); i++ {
		fmt.Printf("%s:", wData[i].key)

		if wData[i].check {
			fmt.Printf("%g\n", wData[i].value)
		} else {
			fmt.Printf("NULL\n")
		}
	}
}

func Clear(wData []Weather) []Weather {
	for i := 0; i < len(wData); i++ {
		wData[i].check = false
	}
	return wData
}
