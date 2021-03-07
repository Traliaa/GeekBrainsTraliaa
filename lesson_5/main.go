package main

import "fmt"

var a int

func main() {

	for {
		fmt.Println("Введите чило")
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Printf("error => %s", err)
		}
		fmt.Println(fib(a))
	}

}
func fib(n int) int {
	numberFib := make(map[int]int)
	numberFib[0] = 0
	numberFib[1] = 1
	if val, ok := numberFib[n]; ok {
		return val
	}
	numberFib[n] = fib(n-1) + fib(n-2)
	return numberFib[n]

}
