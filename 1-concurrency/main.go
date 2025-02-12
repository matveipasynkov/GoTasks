package main

import (
	"fmt"
	"math/rand/v2"
)

func makeRandomNumbers(ch chan int) {
	numbers := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.IntN(101))
	}
	for i := 0; i < len(numbers); i++ {
		ch <- numbers[i]
	}
}

func updateNumbers(chRandom chan int, chMain chan int) {
	var number int;
	for i := 0; i < 10; i++ {
		number = <-chRandom
		chMain <- number * number
	}
}

func main() {
	chFunc := make(chan int, 10)
	chMain := make(chan int, 10)

	go makeRandomNumbers(chFunc)
	go updateNumbers(chFunc, chMain)

	for i := 0; i < 10; i++ {
		fmt.Print(<-chMain, " ")
	}
	close(chFunc)
	close(chMain)
}