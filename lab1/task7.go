package main

import "fmt"

func main() {
	// Оголошуємо змінні типів int8 і int16
	variable8 := int8(127)    // Максимальне значення для типу int8: 127 (01111111 у двійковій формі)
	variable16 := int16(16383) // Число в межах int16: 16383 (0011111111111111 у двійковій формі)

	fmt.Println("Приведення типів\n")

	// Вивід значення variable8 у десятковій і двійковій формах
	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)

	// Вивід значення variable16 у десятковій і двійковій формах
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)

	// Приведення variable8 (int8) до uint16
	// Значення 127 лишається тим самим, але тип змінюється
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))

	// Приведення variable16 (int16) до uint8
	// Оскільки максимальне значення для uint8 - 255, значення variable16 "обрізається" до молодших 8 бітів
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання:
	// Створюємо 2 змінні різних типів та виконуємо арифметичні операції

	// Змінні різних типів
	var a int16 = 30000   // Значення в межах int16
	var b uint8 = 200     // Значення в межах uint8

	// Арифметичні операції з приведенням типів
	sum := a + int16(b)      // Приводимо uint8 до int16 для додавання
	difference := a - int16(b)
	product := a * int16(b)
	quotient := a / int16(b)

	// Виведення результатів у консоль
	fmt.Println("\nАрифметичні операції:")
	fmt.Printf("Сума:           %d + %d = %d\n", a, b, sum)
	fmt.Printf("Різниця:        %d - %d = %d\n", a, b, difference)
	fmt.Printf("Добуток:        %d * %d = %d\n", a, b, product)
	fmt.Printf("Частка:         %d / %d = %d\n", a, b, quotient)
}
