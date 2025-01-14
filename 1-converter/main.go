package main

import "fmt"

func main() {
	const usdEur = 0.98
	const usdRub = 103.05
	eurRub := usdRub / usdEur
	fmt.Println(eurRub)
}

func getData() (int, string, string) {
	var (
		value           int
		currency_first  string
		currency_second string
	)
	fmt.Print("Введите первую валюту: ")
	fmt.Scan(&currency_first)
	fmt.Print("Введите вторую валюту: ")
	fmt.Scan(&currency_second)
	fmt.Print("Значение: ")
	fmt.Scan(&value)
	return value, currency_first, currency_second
}

func calculate(value int, currency_first string, currency_second string) {

}
