package main

import "fmt"

func main() {
	value, currencyFirst, currencySecond := getData()
	result := calculate(value, currencyFirst, currencySecond)
	fmt.Println("Результат в", currencySecond, "равен:", result)
}

func getData() (float64, string, string) {
	var (
		value          float64
		currencyFirst  string
		currencySecond string
	)
	currencyFirst = checkFirstCurrency()
	value = checkValue()
	currencySecond = checkSecondCurrency(currencyFirst)
	return value, currencyFirst, currencySecond
}

func calculate(value float64, currencyFirst string, currencySecond string) float64 {
	const usdEur = 0.98
	const usdRub = 103.05
	const eurRub = usdRub / usdEur

	if currencyFirst == "EUR" && currencySecond == "RUB" {
		return value * eurRub
	} else if currencyFirst == "EUR" && currencySecond == "USD" {
		return value / usdEur
	} else if currencyFirst == "USD" && currencySecond == "RUB" {
		return value * usdRub
	} else if currencyFirst == "USD" && currencySecond == "EUR" {
		return value * usdEur
	} else if currencyFirst == "RUB" && currencySecond == "USD" {
		return value / usdRub
	} else {
		return value / eurRub
	}
}

func checkFirstCurrency() string {
	var result string
	for {
		fmt.Print("Введите первую валюту (EUR, USD, RUB): ")
		fmt.Scan(&result)
		if result == "EUR" || result == "USD" || result == "RUB" {
			return result
		}
		fmt.Println("Валюта введена неверно. Попробуйте ещё раз.")
	}
}

func checkSecondCurrency(firstCurrency string) string {
	var result string
	for {
		fmt.Print("Введите вторую валюту (EUR, USD, RUB): ")
		fmt.Scan(&result)
		if result == "EUR" || result == "USD" || result == "RUB" {
			if firstCurrency != result {
				return result
			}
		}
		fmt.Println("Валюта введена неверно, либо она повторилась. Попробуйте ещё раз.")
	}
}

func checkValue() float64 {
	var result float64
	var discard string
	for {
		fmt.Print("Введите значение (больше 0): ")
		_, err := fmt.Scan(&result)
		if err == nil && result > 0 {
			return result
		}
		fmt.Println("Значение введено неверно. Попробуйте ещё раз.")
		if err == nil {
			continue
		}
		fmt.Scanln(&discard)
	}
}
