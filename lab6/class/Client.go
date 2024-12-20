package class

import "errors"

// Client представляє клієнта банку з основними атрибутами та методами.
type Client struct {
	name          string  // Ім'я клієнта
	surname       string  // Прізвище клієнта
	accountNumber int     // Номер рахунку клієнта
	cDeposit      float64 // Сума депозиту клієнта
	cCredit       float64 // Сума кредиту клієнта
}

// NewClient створює нового клієнта. Перевіряє валідність введених даних і повертає помилку, якщо вони некоректні.
func NewClient(name, surname string, accountNumber int, cDeposit, cCredit float64) (*Client, error) {
	client := new(Client)
	errorString := ""
	client.name = name
	client.surname = surname

	// Перевірка валідності номера рахунку
	if AmountValidator(accountNumber) {
		client.accountNumber = accountNumber
	} else {
		errorString += "invalid account number; "
	}

	// Перевірка валідності депозиту
	if MoneyValidator(cDeposit) {
		client.cDeposit = cDeposit
	} else {
		errorString += "invalid deposit amount; "
	}

	// Перевірка валідності кредиту
	if MoneyValidator(cCredit) {
		client.cCredit = cCredit
	} else {
		errorString += "invalid credit amount; "
	}

	if errorString == "" {
		return client, nil
	} else {
		return nil, errors.New(errorString)
	}
}

// SetName встановлює ім'я клієнта.
func (client *Client) SetName(name string) {
	client.name = name
}

// Name повертає ім'я клієнта.
func (client *Client) Name() string {
	return client.name
}

// SetSurname встановлює прізвище клієнта.
func (client *Client) SetSurname(surname string) {
	client.surname = surname
}

// Surname повертає прізвище клієнта.
func (client *Client) Surname() string {
	return client.surname
}

// SetAccountNumber встановлює номер рахунку клієнта, якщо значення валідне.
func (client *Client) SetAccountNumber(accNumber int) error {
	if AmountValidator(accNumber) {
		client.accountNumber = accNumber
		return nil
	} else {
		return errors.New("invalid account number")
	}
}

// AccountNumber повертає номер рахунку клієнта.
func (client *Client) AccountNumber() int {
	return client.accountNumber
}

// SetCDeposit встановлює суму депозиту клієнта, якщо значення валідне.
func (client *Client) SetCDeposit(cDeposit float64) error {
	if MoneyValidator(cDeposit) {
		client.cDeposit = cDeposit
		return nil
	} else {
		return errors.New("invalid deposit money amount")
	}
}

// CDeposit повертає суму депозиту клієнта.
func (client *Client) CDeposit() float64 {
	return client.cDeposit
}

// SetCCredit встановлює суму кредиту клієнта, якщо значення валідне.
func (client *Client) SetCCredit(cCredit float64) error {
	if MoneyValidator(cCredit) {
		client.cCredit = cCredit
		return nil
	} else {
		return errors.New("invalid credit money amount")
	}
}

// CCredit повертає суму кредиту клієнта.
func (client *Client) CCredit() float64 {
	return client.cCredit
}

// DepositMoneyChange змінює суму депозиту клієнта та оновлює стан банку.
func (client *Client) DepositMoneyChange(bank *Bank, money float64) {
	if money < 0 {
		// Зняття грошей з депозиту
		if bank.BankMoney() < -money {
			return
		}
		client.cDeposit += money
		if client.cDeposit < 0 {
			money -= client.cDeposit
			client.cDeposit = 0
			bank.SetDeposit(bank.Deposit() - 1)
		}
		bank.SetBankMoney(bank.BankMoney() + money)
	} else {
		// Додавання грошей до депозиту
		client.cDeposit += money
		bank.SetBankMoney(bank.BankMoney() + money)
	}
}

// CreditMoneyChange змінює суму кредиту клієнта та оновлює стан банку.
func (client *Client) CreditMoneyChange(bank *Bank, money float64) {
	if money > 0 {
		// Отримання кредиту
		if bank.BankMoney() < money {
			return
		}
		client.cCredit += money
		bank.SetBankMoney(bank.BankMoney() - money)
	} else {
		// Погашення кредиту
		client.cCredit += money
		if client.cCredit < 0 {
			money -= client.cCredit
			client.cCredit = 0
			bank.SetCredit(bank.Credit() - 1)
		}
		bank.SetBankMoney(bank.BankMoney() - money)
	}
}
