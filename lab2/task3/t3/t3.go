package t3

// Функція для знаходження мінімального значення серед трьох чисел
func findMinimum(val1, val2, val3 float32) float32 {
    var minimum float32

    // Порівнюємо три числа і визначаємо мінімальне
    if val1 < val2 && val1 < val3 {
        minimum = val1
    } else if val2 < val1 && val2 < val3 {
        minimum = val2
    } else {
        minimum = val3
    }

    return minimum  // Повертаємо мінімальне значення
}

// Функція для обчислення середнього значення трьох чисел
func calculateAverage(num1, num2, num3 float32) float32 {
    return (num1 + num2 + num3) / 3  // Повертаємо середнє значення
}

// Функція для обчислення результату рівняння першого порядку: y = kx + b
func firstOrderEquation(k, x, b float32) float32 {
    return k*x + b  // Повертаємо результат рівняння
}
