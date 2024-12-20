package class

// MoneyValidator перевіряє, чи є сума грошей валідною (не від'ємною).
func MoneyValidator(money float64) bool {
	return money >= 0.0
}

// AmountValidator перевіряє, чи є кількість валідною (не від'ємною).
func AmountValidator(amount int) bool {
	return amount >= 0
}

// Унікальний ідентифікатор для клієнтів банку.
var id int = 0

// GetClientId генерує та повертає новий унікальний ідентифікатор для клієнта.
func GetClientId() int {
	id++
	return id
}
