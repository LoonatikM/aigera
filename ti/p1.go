package main

import (
	"fmt"
	"math"
)

func main() {
	go6()
}

func go1() {
	// Обмен значений a и b без третьей переменной
	a, b := 10, 6
	a, b = b, a
	fmt.Println(a, b)
}

func go2() {
	// Конвертер Цельсий -> Фаренгейт.
	var a float64
	fmt.Scan(&a)
	b := a*9/5 + 32
	fmt.Println("result: ", b)
}

func go3() {
	// Площадь и периметр прямоугольника.
	var a int
	var b int
	fmt.Printf("Enter a: ")
	fmt.Scan(&a)
	fmt.Printf("Enter b: ")
	fmt.Scan(&b)

	s := a * b
	p := 2 * (a + b)

	fmt.Println("Result s: ", s, "and p: ", p)
}

func go4() {
	// Расчет гипотенузы катетов.
	var a, b, x float64
	fmt.Printf("enter a: ")
	fmt.Scan(&a)
	fmt.Printf("enter b: ")
	fmt.Scan(&b)
	x = math.Pow(a, 2) + math.Pow(b, 2)
	fmt.Println("Result: ", math.Sqrt(x))
}

func go5() {
	// Среднее арифметическое трех чисел.
	var a, b, c, x float64
	fmt.Printf("Enter a: ")
	fmt.Scan(&a)
	fmt.Printf("Enter b: ")
	fmt.Scan(&b)
	fmt.Printf("Enter c: ")
	fmt.Scan(&c)
	x = (a + b + c) / 3
	fmt.Println("Result: ", x)
}

func go6() {
	// Вычисление остатка от деления без %.
	var a, b, c, d int
	a = 200
	b = 35
	d = a % b
	c = a - ((a / b) * b)
	fmt.Println("Result c go6: ", c)
	fmt.Println("Result d go6: ", d)
}
