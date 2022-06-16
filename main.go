package main

import (
	"fmt"
)

func main() {
	var our_money float32
	const apple_price = 5.99
	const pear_price = 7.0
	our_money = 23.0
	fmt.Println("Скільки треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println(9*apple_price + 8*pear_price)
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Println(int(our_money / pear_price))
	fmt.Println("Скільки тяблук ми можемо купити?")
	fmt.Println(int(our_money / apple_price))
	fmt.Println("Чм можемо ми купити 2 груші та 2 яблука?")
	fmt.Println((2*pear_price + 2*apple_price) <= our_money)
}
