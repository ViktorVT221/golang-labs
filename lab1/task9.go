package main

import "fmt"

func main() {
	// Оголошення булевих змінних зі значеннями за замовчуванням
	var first, second bool  // Значення за замовчуванням для bool: false
	var third bool = true   // Явна ініціалізація значення true
	fourth := !third        // Заперечення значення third: !true → false
	var fifth = true        // Явна ініціалізація значення true

	// Виведення значень булевих змінних
	fmt.Println("first  = ", first)       // false: значення за замовчуванням
	fmt.Println("second = ", second)      // false: значення за замовчуванням
	fmt.Println("third  = ", third)       // true: було ініціалізовано як true
	fmt.Println("fourth = ", fourth)      // false: !third інвертує значення (true → false)
	fmt.Println("fifth  = ", fifth, "\n") // true: було ініціалізовано як true

	// Логічна операція "НЕ" (!)
	fmt.Println("!true  = ", !true)        // false: заперечення true
	fmt.Println("!false = ", !false, "\n") // true: заперечення false

	// Логічна операція "І" (&&)
	fmt.Println("true && true   = ", true && true)         // true: обидва операнди true
	fmt.Println("true && false  = ", true && false)        // false: один із операндів false
	fmt.Println("false && false = ", false && false, "\n") // false: обидва операнди false

	// Логічна операція "АБО" (||)
	fmt.Println("true || true   = ", true || true)         // true: хоча б один операнд true
	fmt.Println("true || false  = ", true || false)        // true: хоча б один операнд true
	fmt.Println("false || false = ", false || false, "\n") // false: обидва операнди false

	// Порівняння чисел
	fmt.Println("2 < 3  = ", 2 < 3)        // true: 2 менше за 3
	fmt.Println("2 > 3  = ", 2 > 3)        // false: 2 не більше за 3
	fmt.Println("3 < 3  = ", 3 < 3)        // false: 3 не менше за 3
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true: 3 менше або дорівнює 3
	fmt.Println("3 > 3  = ", 3 > 3)        // false: 3 не більше за 3
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true: 3 більше або дорівнює 3
	fmt.Println("2 == 3 = ", 2 == 3)       // false: 2 не дорівнює 3
	fmt.Println("3 == 3 = ", 3 == 3)       // true: 3 дорівнює 3
	fmt.Println("2 != 3 = ", 2 != 3)       // true: 2 не дорівнює 3
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false: 3 дорівнює 3, тому результат false
}
