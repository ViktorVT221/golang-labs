package main

import "fmt"

func main() {
    // Обчислюємо мінімальне значення з трьох заданих
    resultMin := getMinValue(0.1, -7, 2)
    fmt.Println("Мінімальне значення:", resultMin)

    // Обчислюємо середнє значення трьох чисел
    avg := calculateMean(14, 102, 501)
    fmt.Println("Середнє значення:", avg)

    // Розв'язуємо рівняння першого порядку y = kx + b
    equation := firstOrderEquation(6, 10, 5)
    fmt.Println("Рівняння першого порядку: y=kx+b\nРозв'язок:", equation)
}

// Функція для знаходження мінімального значення серед трьох заданих
// Порівнює значення та повертає найменше
func getMinValue(val1, val2, val3 float32) float32 {
    // Ініціалізуємо змінну для мінімального значення
    var minVal float32

    // Порівнюємо значення та визначаємо найменше
    if val1 < val2 && val1 < val3 {
        minVal = val1
    } else if val2 < val1 && val2 < val3 {
        minVal = val2
    } else {
        minVal = val3
    }

    // Повертаємо мінімальне значення
    return minVal
}

// Функція для обчислення середнього значення трьох чисел
// Сума чисел ділиться на кількість елементів (3)
func calculateMean(num1, num2, num3 float32) float32 {
    // Обчислюємо середнє значення
    return (num1 + num2 + num3) / 3
}

// Функція для розв'язку рівняння першого порядку: y = kx + b
// Розв'язок обчислюється за формулою
func firstOrderEquation(k, x, b float32) float32 {
    // Обчислюємо значення y
    return k*x + b
}
