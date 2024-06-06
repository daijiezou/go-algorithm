package main

import "fmt"

func main() {
	var high float64
	fmt.Scan(&high)
	totalDistance := high
	for i := 0; i < 5; i++ {
		high = high / 2
		totalDistance += high * 2
	}
	fmt.Println(totalDistance - high*2)
	fmt.Println(high)
}
