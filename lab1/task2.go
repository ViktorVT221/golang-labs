package main

import "fmt"

func main() {
	// Усі змінні ініціалізуються значенням за замовчуванням, тобто 0
	var defaultInt8 int8
	var defaultInt16 int16
	var defaultInt32 int32
	var defaultInt64 int64
	var defaultInt int

	// Виведення значень за замовчуванням для цілочисельних типів (зі знаком)
	fmt.Println("Значення за замовчуванням (зі знаком):", defaultInt8, defaultInt16, defaultInt32, defaultInt64, defaultInt)

	// Аналогічно для беззнакових типів
	var defaultUInt8 uint8
	var defaultUInt16 uint16
	var defaultUInt32 uint32
	var defaultUInt64 uint64
	var defaultUInt uint

	// Виведення значень за замовчуванням для беззнакових типів
	fmt.Println("Значення за замовчуванням (без знака):", defaultUInt8, defaultUInt16, defaultUInt32, defaultUInt64, defaultUInt)

	// Завдання:
	// 1. Створити цілочисельну змінну (результат не відображати)
	// Наприклад:
	var myInt8 int8 = 8  // Цілочисельна змінна типу int8 зі значенням 8
	var myInt16 int16 = 16 // Цілочисельна змінна типу int16 зі значенням 16
	var myInt32 int32 = 32 // Цілочисельна змінна типу int32 зі значенням 32
	var myInt64 int64 = 64 // Цілочисельна змінна типу int64 зі значенням 64
	var myInt int = 1      // Цілочисельна змінна типу int зі значенням 1

	// Ці змінні оголошені, але їх значення не виводяться на екран
	_ = myInt8
	_ = myInt16
	_ = myInt32
	_ = myInt64
	_ = myInt
}
