package main

import (
	"fmt"
	"math"
)

// 9 варіант
func main() {
	var x0 int32 = 0
	var arr []int32
	GenerateRandomValues(x0, &arr) // Генерація випадкових значень
	ReverseArray(&arr)             // Перевернути масив

	// Створення масиву частот
	var freq [250]int32
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	// Виведення статистики
	for i := 0; i < 250; i++ {
		if freq[i] > 0 {
			probability := float32(freq[i]) / 20000
			fmt.Printf("№ %d %d К-ть повторів: %d Статична ймовірність: %.6f\n", i+1, i, freq[i], probability)
		}
	}

	// Математичне сподівання
	exp := CalculateExpectation(&arr, &freq)
	fmt.Printf("\nМатематичне сподівання: %.6f", exp)

	// Обчислення дисперсії
	var variance float64
	for i := 0; i < 250; i++ {
		if freq[i] > 0 {
			variance += math.Pow(float64(i)-exp, 2) * float64(freq[i]) / 20000
		}
	}
	fmt.Printf("\nДисперсія: %.6f", variance)

	// Середньоквадратичне відхилення
	sigma := CalculateSigma(&arr)
	fmt.Printf("\nСередньоквадратичне відхилення: %.6f", sigma)
}

// Генерація випадкових чисел
func GenerateRandomValues(x0 int32, arr *[]int32) {
	const a = 1140671485
	const c = 12820163
	const m = 16777216 // 2^24

	// Генерація 20000 випадкових чисел
	for k := 0; k < 20000; k++ {
		x0 = (a*x0 + c) % m
		*arr = append(*arr, x0%250) // Генерація числа від 0 до 249
	}
}

// Обчислення математичного сподівання
func CalculateExpectation(arr *[]int32, frequency *[250]int32) float64 {
	var expectation float64
	for i := 0; i < len(*arr); i++ {
		expectation += float64((*arr)[i]) * float64((*frequency)[(*arr)[i]]) / 20000
	}
	return expectation
}

// Обчислення середньоквадратичного відхилення
func CalculateSigma(arr *[]int32) float64 {
	var sigma float64
	var sum float64
	var n = len(*arr)

	// Обчислення середнього значення
	for i := 0; i < n; i++ {
		sum += float64((*arr)[i])
	}
	avg := sum / float64(n)

	// Обчислення дисперсії
	for i := 0; i < n; i++ {
		sigma += math.Pow(float64((*arr)[i])-avg, 2)
	}
	sigma = math.Sqrt(sigma / float64(n))
	return sigma
}

// Перевернути масив
func ReverseArray(numbers *[]int32) {
	for i := 0; i < len(*numbers)/2; i++ {
		j := len(*numbers) - i - 1
		(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
	}
}
