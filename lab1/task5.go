package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32 або int64, в залежності від платформи")
	fmt.Println("uint    - uint32 або uint64, в залежності від платформи")

	// Визначаємо розмір int у байтах
	var x int
	size := unsafe.Sizeof(x) // unsafe.Sizeof повертає розмір типу у байтах

	// Визначаємо розрядність на основі розміру
	if size == 4 {
		fmt.Println("\nПлатформа є 32-бітною")
	} else if size == 8 {
		fmt.Println("\nПлатформа є 64-бітною")
	} else {
		fmt.Println("\nНевідома розрядність платформи")
	}
}
