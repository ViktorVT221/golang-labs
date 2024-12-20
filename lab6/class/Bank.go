package class

import "errors"

// Bank представляє банк з основними атрибутами та методами.
type Bank struct {
	name          string    // Назва банку
	bankMoney     float64   // Загальна сума грошей у банку
	deposit       int       // Депозитна ставка (у відсотках)
	credit        int       // Кредитна ставка (у відсотках)
	clientsAmount int       // Кількість клієнтів банку
	clients       []*Client // Список клієнтів банку
}

// NewBank створює новий банк. Перевіряє валідність введених даних і повертає помилку, якщо вони некоректні.
func NewBank(name string, bankMoney float64, deposit, credit int, clients []*Client) (*Bank, error) {
	bank := new(Bank)
	errorString := ""
	bank.name = name

	// Перевірка валідності суми грошей
	if MoneyValidator(bankMoney) {
		bank.bankMoney = bankMoney
	} else {
		errorString += "invalid money amount; "
	}

	// Перевірка валідності депозитної ставки
	if AmountValidator(deposit) {
		bank.deposit = deposit
	} else {
		errorString += "invalid deposit amount; "
	}

	// Перевірка валідності кредитної ставки
	if AmountValidator(credit) {
		bank.credit = credit
	} else {
		errorString += "invalid credit amount; "
	}

	bank.clients = clients

	if errorString == "" {
		return bank, nil
	} else {
		return nil, errors.New(errorString)
	}
}

// SetName встановлює назву банку.
func (bank *Bank) SetName(name string) {
	bank.name = name
}

// Name повертає назву банку.
func (bank *Bank) Name() string {
	return bank.name
}

// SetBankMoney встановлює суму грошей у банку, якщо значення валідне.
func (bank *Bank) SetBankMoney(money float64) error {
	if MoneyValidator(money) {
		bank.bankMoney = money
		return nil
	} else {
		return errors.New("invalid money amount")
	}
}

// BankMoney повертає суму грошей у банку.
func (bank *Bank) BankMoney() float64 {
	return bank.bankMoney
}

// SetDeposit встановлює депозитну ставку, якщо значення валідне.
func (bank *Bank) SetDeposit(deposit int) error {
	if AmountValidator(deposit) {
		bank.deposit = deposit
		return nil
	} else {
		return errors.New("invalid deposit amount")
	}
}

// Deposit повертає депозитну ставку.
func (bank *Bank) Deposit() int {
	return bank.deposit
}

// SetCredit встановлює кредитну ставку, якщо значення валідне.
func (bank *Bank) SetCredit(credit int) error {
	if AmountValidator(credit) {
		bank.credit = credit
		return nil
	} else {
		return errors.New("invalid credit amount")
	}
}

// Credit повертає кредитну ставку.
func (bank *Bank) Credit() int {
	return bank.credit
}

// SetClientsAmount встановлює кількість клієнтів, якщо значення валідне.
func (bank *Bank) SetClientsAmount(clientsAmount int) error {
	if AmountValidator(clientsAmount) {
		bank.clientsAmount = clientsAmount
		return nil
	} else {
		return errors.New("invalid clients amount")
	}
}

// ClientsAmount повертає кількість клієнтів банку.
func (bank *Bank) ClientsAmount() int {
	return bank.clientsAmount
}

// SetClients встановлює список клієнтів банку.
func (bank *Bank) SetClients(clients []*Client) {
	bank.clients = clients
}

// Clients повертає список клієнтів банку.
func (bank *Bank) Clients() []*Client {
	return bank.clients
}
