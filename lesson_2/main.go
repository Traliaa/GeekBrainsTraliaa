package main

import (
	"fmt"
	"math"
)

func main() {
	//rectangle()
	//rectangleСircle()
	number()
}

func rectangle() {
	var a, b int
	fmt.Println("Введите длину первый стороны")
	fmt.Scan(&a)
	fmt.Println("Введите длину второй стороны")
	fmt.Scan(&b)
	fmt.Printf("Площадь прямоугольника %d", a*b)
}

//Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
func rectangleСircle() {
	var a float64
	fmt.Println("Введите площадь круга")
	fmt.Scan(&a)
	b := 2 * math.Sqrt(a/math.Pi)
	c := math.Sqrt(a * 4 * math.Pi)
	fmt.Printf("Диамтер круга =  %0.2f, длина окружности %0.2f", b, c)
}

//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе
func number() {
	var x int
	fmt.Println("Введите число")
	fmt.Scan(&x)
	a := x / 100
	b := x % 100 / 10
	c := x % 10
	fmt.Printf("Сотен = %d, десятков =%d, едениц=%d", a, b, c)
}
