package main

import (
	"fmt"
	class "Lab6/class"       // Імпортуємо модуль класів.
	mfunc "Lab6/mainfunctions" // Імпортуємо модуль основних функцій.
)

func main() {
	// Ініціалізація масиву банків і допоміжних змінних.
	banks := make([]*class.Bank, 10) // Масив банків.
	banksAmount := 0                // Кількість банків.
	idClient := 1                   // Початковий ідентифікатор клієнта.
	action := -1                    // Дія, обрана користувачем.
	check := false                  // Прапорець для завершення програми.

	// Головний цикл програми.
	for !check {
		// Вивід меню для вибору дії.
		fmt.Println("\nChoose your action: ")
		fmt.Println("0 - exit")
		fmt.Println("1 - create bank")
		fmt.Println("2 - add client for deposit")
		fmt.Println("3 - add client for credit")
		fmt.Println("4 - client info by surname")
		fmt.Println("5 - client info by id")
		if action != -1 {
			fmt.Println(banks[0]) // Вивід інформації про перший банк (для перевірки).
		}
		fmt.Scanf("%d\n", &action) // Зчитуємо вибір користувача.

		// Обробка вибору дії.
		switch action {
		case 1:
			{
				// Додавання нового банку.
				banks[banksAmount] = mfunc.AddBank()
				banksAmount++
				fmt.Scanf("\n")
			}
		case 2:
			{
				// Додавання клієнта для депозиту.
				err := mfunc.AddDepositClient(banks, banksAmount, idClient)
				if err == nil {
					idClient++ // Збільшуємо ID клієнта.
					fmt.Println("Client added")
				} else {
					fmt.Println(err)
				}
				fmt.Scanf("\n")
			}
		case 3:
			{
				// Додавання клієнта для кредиту.
				err := mfunc.AddCreditClient(banks, banksAmount, idClient)
				if err == nil {
					idClient++ // Збільшуємо ID клієнта.
					fmt.Println("Client added")
				} else {
					fmt.Println(err)
				}
				fmt.Scanf("\n")
			}
		case 4:
			{
				// Пошук клієнта за прізвищем.
				client, err := mfunc.FindClientBySurname(banks, banksAmount)
				if err == nil {
					mfunc.PrintClientInfo(client)
				} else {
					fmt.Println(err)
				}
				fmt.Scanf("\n")
			}
		case 5:
			{
				// Пошук клієнта за ID.
				client, err := mfunc.FindClientById(banks, banksAmount)
				if err == nil {
					mfunc.PrintClientInfo(client)
				} else {
					fmt.Println(err)
				}
				fmt.Scanf("\n")
			}
		case 0:
			{
				// Завершення роботи програми.
				check = true
				fmt.Println("Ending work...")
			}
		default:
			{
				// Обробка некоректного вибору.
				fmt.Println("invalid action")
			}
		}
	}
}
