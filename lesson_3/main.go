package main

import (
	"fmt"
	"log"
)

func main() {
	calculator()
}

func calculator() {
	var a, b, result float64
	var operator string

	fmt.Println("Введите первое чило")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal("Введено неверное число!")
	}

	_, err = fmt.Println("Введите второе число")
	if err != nil {
		log.Fatal("Введено неверное число!")
	}
	fmt.Scan(&b)
	fmt.Println("Введите операцию которую необходимо сделать")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatal("Введено неверный опертаор!")
	}

	switch operator {
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "+":
		result = a + b
	case "-":
		result = a * b
	}
	fmt.Println(result)

}
