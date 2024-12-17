package main

// Імпорт кількох стандартних пакетів
import (
	"fmt"  // Пакет для форматованого виводу
	"math" // Пакет для роботи з математичними константами і функціями
)

func main() {
	// Оголошення змінних з ініціалізацією за замовчуванням
	var defaultFloat float32       // Значення за замовчуванням для float32 = 0
	var defaultDouble float64 = 5.5 // Явна ініціалізація для float64

	// Виведення змінних і їх типів
	fmt.Println("defaultFloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	// Максимальні та мінімальні значення float32
	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	// Максимальні та мінімальні значення float64
	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання:
	// Створення змінних різних типів за допомогою короткого запису
	intVar := 42          // int за замовчуванням
	stringVar := "Hello!" // string
	boolVar := true       // bool
	floatVar := 3.14      // float64
	byteVar := byte(255)  // byte (uint8)

	// Виведення значень і типів змінних
	fmt.Println("Змінні, створені коротким записом:")
	fmt.Printf("intVar    (%T) = %d\n", intVar, intVar)
	fmt.Printf("stringVar (%T) = %s\n", stringVar, stringVar)
	fmt.Printf("boolVar   (%T) = %t\n", boolVar, boolVar)
	fmt.Printf("floatVar  (%T) = %.2f\n", floatVar, floatVar)
	fmt.Printf("byteVar   (%T) = %d\n", byteVar, byteVar)
}
