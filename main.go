package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("\033[35;1mWELCOME TO UNIT CONVERTER ⚒︎ \033[0m\n")
	reader := bufio.NewReader(os.Stdin)

	conversionOptions := map[int]string{
		1: "METERS to KILOMETERS",
	}

	for {
		printConversionOptions(conversionOptions)
		fmt.Print("\033[33;1m Enter your CHOICE: \033[0m")

		choice, err := readInt(reader)

		if err != nil {
			fmt.Println("\n\033[31;1m✘ INVALID INPUT, PLEASE ENTER A VALID NUMBER.\033[0m")
			continue
		}

		if choice == 0 {
			fmt.Println("\n\033[32;1mEXITING THE PROGRAM. GOODBYE! ⚒︎\033[0m")
			break
		}
	}
}

func printConversionOptions(options map[int]string) {
	fmt.Println("\n\033[33;1m CHOOSE A CONVERSION: \033[0m")
	for key, value := range options {
		fmt.Printf("%d. %s\n", key, value)
	}
	fmt.Println("0. PROGRAM EXIT")
}

func readInt(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}
