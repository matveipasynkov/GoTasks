package main

import "fmt"

func main() {
	const usdEur = 0.98
	const usdRub = 103.05
	eurRub := usdRub / usdEur
	fmt.Println(eurRub)
}
