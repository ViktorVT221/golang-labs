package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1          // Змінна типу uint8
	var userinit16 uint16 = 2        // Змінна типу uint16
	var userinit64 int64 = -3        // Змінна типу int64
	var userautoinit = -4            // Автоматичне визначення типу (тут int)

	// Вивід значень змінних
	fmt.Println("Values:", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// Тип задається автоматично за значенням
	intVar := 10

	// Вивід значення та типу змінної
	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання 1: Вивести типи всіх змінних
	fmt.Printf("Type of userinit8 = %T\n", userinit8)
	fmt.Printf("Type of userinit16 = %T\n", userinit16)
	fmt.Printf("Type of userinit64 = %T\n", userinit64)
	fmt.Printf("Type of userautoinit = %T\n\n", userautoinit)

	// Завдання 2: Присвоїти змінній intVar значення userinit16 та userautoinit

	// Присвоїти userinit16 (виконається явне преведення типів)
	intVar = int(userinit16)
	fmt.Printf("After assigning userinit16: Value = %d\n", intVar)

	// Присвоїти userautoinit (типи співпадають, тому пряме присвоєння)
	intVar = userautoinit
	fmt.Printf("After assigning userautoinit: Value = %d\n", intVar)
}
