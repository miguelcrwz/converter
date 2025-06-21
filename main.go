package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/miguelcrwz/conversions"
)

func main() {
	fmt.Printf("\033[35;1mWELCOME TO UNIT CONVERTER ⚒︎ \033[0m\n")
	reader := bufio.NewReader(os.Stdin)

	conversionOptions := map[int]string{
		1: "METERS to KILOMETERS",
		2: "KILOGRAMS to POUNDS",
		3: "CELSIUS to FAHRENHEIT",
		4: "SECONDS to MINUTES",
		5: "SQUARE METERS to SQUARE KILOMETERS",
		6: "LITERS to GALLONS",
	}

	for {
		printConversionOptions(conversionOptions)
		fmt.Print("\033[33;1m Enter your CHOICE: \033[0m")

		choice, err := readInt(reader)

		if err != nil {
			pause()
			fmt.Println("\n\033[31;1m✘ INVALID INPUT, PLEASE ENTER A VALID NUMBER.\033[0m")
			continue
		}

		if choice == 0 {
			pause()
			fmt.Println("\n\033[32;1mEXITING THE PROGRAM. GOODBYE! ⚒︎\033[0m")
			break
		}

		if _, exists := conversionOptions[choice]; !exists {
			pause()
			fmt.Println("\n\033[31;1m✘ INVALID CHOICE, PLEASE CHOOSE A VALID OPTION.\033[0m")
			continue
		}

		handleConversion(choice, reader)
	}
}

func printConversionOptions(options map[int]string) {
	fmt.Println("\n\033[33;1m CHOOSE A CONVERSION: \033[0m")
	pause()
	for key, value := range options {
		fmt.Printf("%d. %s\n", key, value)
		pause()
	}
	fmt.Println("0. PROGRAM EXIT")
	pause()
}

func handleConversion(choice int, reader *bufio.Reader) {
	switch choice {
	case 1:
		value := readFloat(reader, "\n\033[35;1mEnter VALUE in METERS: \033[0m")
		pause()
		fmt.Printf("\033[35;1m%.2f METERS = %.3f KILOMETERS\033[0m\n", value, conversions.MetersToKilometers(value))
	case 2:
		value := readFloat(reader, "\n\033[35;1mEnter VALUE in KILOGRAMS: \033[0m")
		pause()
		fmt.Printf("\n\033[35;1m%.2f KILOGRAMS = %.3f POUNDS\033[0m\n", value, conversions.KilogramsToPounds(value))
	case 3:
		value := readFloat(reader, "\n\033[35;1mEnter TEMPERATURE in CELSIUS: \033[0m")
		pause()
		fmt.Printf("\n\033[35;1m%.2f CELSIUS = %.3f FAHRENHEIT\033[0m\n", value, conversions.CelsiusToFahrenheit(value))
	case 4:
		value := readFloat(reader, "\n\033[35;1mEnter TIME in SECONDS: \033[0m")
		pause()
		fmt.Printf("\n\033[35;1m%.2f SECONDS = %.3f MINUTES\033[0m\n", value, conversions.SecondsToMinutes(value))
	case 5:
		value := readFloat(reader, "\n\033[35;1mEnter AREA in SQUARE METERS: \033[0m")
		pause()
		fmt.Printf("\n\033[35;1m%.2f SQUARE METERS = %.3f SQUARE KILOMETERS\033[0m\n", value, conversions.SquareMetersToSquareKilometers(value))
	case 6:
		value := readFloat(reader, "\n\033[35;1mEnter VOLUME in LITERS: \033[0m")
		pause()
		fmt.Printf("\n\033[35;1m%.2f LITERS = %.3f GALLONS\033[0m\n", value, conversions.LitersToGallons(value))
	}
}

func readInt(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			pause()
			fmt.Println("\n\033[31;1m✘ INVALID INPUT, PLEASE ENTER A VALID NUMBER.\033[0m")
			continue
		}
		return value
	}
}

func pause() {
	time.Sleep(time.Millisecond * 370)
}
