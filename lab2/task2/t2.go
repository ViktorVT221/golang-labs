package main

import "fmt"

// Функція для знаходження мінімального значення серед трьох
func findMinValue(val1, val2, val3 float32) float32 {
    var minValue float32

    if val1 < val2 && val1 < val3 {
        minValue = val1
    } else if val2 < val1 && val2 < val3 {
        minValue = val2
    } else {
        minValue = val3
    }

    return minValue
}

// Функція для обчислення середнього значення трьох чисел
func calculateAverage(num1, num2, num3 float32) float32 {
    return (num1 + num2 + num3) / 3
}

// Функція для обчислення рівняння першого порядку: y = kx + b
func firstOrderEquation(k, x, b float32) float32 {
    return k * x + b
}

func main() {
    // Обчислюємо мінімальне значення серед трьох чисел
    min := findMinValue(0.1, -7, 2)
    fmt.Printf("Мінімальне значення: %.2f\n", min)

    // Обчислюємо середнє значення
    average := calculateAverage(14, 102, 501)
    fmt.Printf("Середнє значення: %.2f\n", average)

    // Обчислюємо результат рівняння першого порядку
    result := firstOrderEquation(6, 10, 5)
    fmt.Printf("Рівняння першого порядку: y = kx + b\nРозв'язок: %.2f\n", result)
}
