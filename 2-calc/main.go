package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	command := getCommand()
	numbers := getNumbers()
	calculation := calculateFunc[command](numbers...)
	fmt.Printf("Результат: %.2f\n", calculation)
}

var calculateFunc = map[string]func(...float64)float64 {
	"MED": median,
	"SUM": sum,
	"AVG": average, 
}

func median(numbers ...float64) float64 {
	if len(numbers)%2 == 0 {
		return (numbers[(len(numbers)-1)/2] + numbers[(len(numbers)-1)/2+1]) / 2
	} else {
		return numbers[(len(numbers)-1)/2]
	}
}

func sum(numbers ...float64) float64 {
	var summary float64 = 0
	for _, value := range numbers {
		summary += value
	}
	return summary
}

func average(numbers ...float64) float64 {
	var summary float64 = 0
	for _, value := range numbers {
		summary += value
	}
	return summary / float64(len(numbers))
}

func getCommand() string {
	var command string
	for {
		fmt.Print("Введите команду (MED, AVG, SUM): ")
		fmt.Scan(&command)
		if command == "MED" || command == "AVG" || command == "SUM" {
			return command
		}
		fmt.Println("Команда неверная. Попробуйте ещё раз.")
	}
}

func getNumbers() []float64 {
	var (
		stringOfNumbers string
		numbers         []float64 = []float64{}
	)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите числа через запятую: ")
		stringOfNumbers, _ = reader.ReadString('\n')
		sliceOfStrings := strings.Split(stringOfNumbers, ",")
		broken := false
		for _, value := range sliceOfStrings {
			floatValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err != nil {
				fmt.Println("Числа введены неверно, попробуйте ещё раз.")
				broken = true
				break
			}
			numbers = append(numbers, floatValue)
		}
		if len(numbers) == 0 || broken {
			if len(numbers) == 0 {
				fmt.Println("Список пустой. Попробуйте ещё раз.")
			}
			continue
		}
		sort.Slice(numbers, func(i int, j int) bool {
			return numbers[i] < numbers[j]
		})
		return numbers
	}
}
