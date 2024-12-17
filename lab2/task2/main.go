package main

import "fmt"

func main() {
	// Обчислюємо мінімальне значення
	var min = getMinValue(0.1, -7, 2)
	fmt.Println("Мінімальне значення:", min)

	// Обчислюємо середнє значення
	var mean = calculateMean(14, 102, 501)
	fmt.Println("Середнє значення:", mean)

	// Розв'язуємо рівняння першого порядку
	var eq = firstOrderEquation(6, 10, 5)
	fmt.Println("Рівняння першого порядку: y=kx+b. \ny=", eq)
}

// Функція для знаходження мінімального значення серед трьох заданих
func getMinValue(val1, val2, val3 float32) float32 {
	if val1 < val2 && val1 < val3 {
		return val1
	} else if val2 < val1 && val2 < val3 {
		return val2
	} else {
		return val3
	}
}

// Функція для обчислення середнього значення трьох чисел
func calculateMean(num1, num2, num3 float32) float32 {
	return (num1 + num2 + num3) / 3
}

// Функція для розв'язку рівняння першого порядку: y = kx + b
func firstOrderEquation(k, x, b float32) float32 {
	return k*x + b
}
