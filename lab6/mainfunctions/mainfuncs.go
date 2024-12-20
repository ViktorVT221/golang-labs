package mainfunctions

import (
	class "Lab6/class"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// AddBank створює новий банк на основі введених користувачем даних.
func AddBank() *class.Bank {
	name := ""
	fmt.Print("Bank name: ")
	fmt.Scanf("%s\n", &name)
	bankMoney := 0.0
	fmt.Print("Bank money: ")
	fmt.Scanf("%f\n", &bankMoney)

	bank, err := class.NewBank(name, bankMoney, 0, 0, make([]*class.Client, 10))
	if err == nil {
		fmt.Println("\nBank added")
	} else {
		fmt.Println(err)
	}
	return bank
}

// AddClient створює нового клієнта на основі введених користувачем даних.
func AddClient(id int) *class.Client {
	name := ""
	fmt.Print("Client name: ")
	fmt.Scanf("%s\n", &name)
	surname := ""
	fmt.Print("Client surname: ")
	fmt.Scanf("%s\n", &surname)

	client, _ := class.NewClient(name, surname, id, 0, 0)
	return client
}

// AddDepositClient додає клієнта з депозитом до вибраного банку.
func AddDepositClient(banks []*class.Bank, banksAmount, id int) error {
	bankName := ""
	fmt.Print("Bank name: ")
	fmt.Scanf("%s\n", &bankName)
	for i := 0; i < banksAmount; i++ {
		bank := banks[i]
		if bank.Name() == bankName {
			client := AddClient(id)

			depositSum := 0.0
			fmt.Print("Enter deposit money: ")
			fmt.Scanf("%f\n", &depositSum)
			err := client.SetCDeposit(depositSum)
			if err != nil {
				return err
			}

			bank.Clients()[bank.ClientsAmount()] = client
			bank.SetClientsAmount(bank.ClientsAmount() + 1)
			bank.SetDeposit(bank.Deposit() + 1)
			bank.SetBankMoney(bank.BankMoney() + depositSum)

			depositRout := make(chan float64, 1)
			exit := make(chan bool)
			depositRout <- depositSum
			go func() {
				for {
					select {
					case <-depositRout:
						{
							for {
								randInput := -depositSum + rand.Float64()*(depositSum-(-depositSum))
								client.DepositMoneyChange(bank, randInput)
								time.Sleep(time.Second * 3)
								if client.CDeposit() <= 0 {
									exit <- true
								}
							}
						}
					case <-exit:
						{
							fmt.Printf("Клієнт %s %s забрав депозит в розмірі %.2f$\n",
								client.Name(), client.Surname(), depositSum)
							return
						}
					}
				}
			}()

			return nil
		}
	}
	return errors.New("this bank doesn't exists")
}

// AddCreditClient додає клієнта з кредитом до вибраного банку.
func AddCreditClient(banks []*class.Bank, banksAmount, id int) error {
	bankName := ""
	fmt.Print("Bank name: ")
	fmt.Scanf("%s\n", &bankName)
	for i := 0; i < banksAmount; i++ {
		bank := banks[i]
		if bank.Name() == bankName {
			client := AddClient(id)

			creditSum := 0.0
			fmt.Print("Enter credit money: ")
			fmt.Scanf("%f\n", &creditSum)
			err := client.SetCCredit(creditSum)
			if err != nil {
				return err
			}
			if bank.BankMoney() < creditSum {
				return errors.New("bank doesn't have enough money")
			}

			bank.Clients()[bank.ClientsAmount()] = client
			bank.SetClientsAmount(bank.ClientsAmount() + 1)
			bank.SetCredit(bank.Credit() + 1)
			bank.SetBankMoney(bank.BankMoney() - creditSum)

			creditRout := make(chan float64, 1)
			exit := make(chan bool)
			creditRout <- creditSum
			go func() {
				for {
					select {
					case <-creditRout:
						{
							for {
								randInput := -creditSum + rand.Float64()*(creditSum-(-creditSum))
								client.CreditMoneyChange(bank, randInput)
								time.Sleep(time.Second * 3)
								if client.CCredit() <= 0 {
									exit <- true
								}
							}
						}
					case <-exit:
						{
							fmt.Printf("Клієнт %s %s виплатив кредит в розмірі %.2f$\n",
								client.Name(), client.Surname(), creditSum)
							return
						}
					}
				}
			}()

			return nil
		}
	}
	return errors.New("this bank doesn't exists")
}

// PrintClientInfo виводить інформацію про клієнта.
func PrintClientInfo(client *class.Client) {
	fmt.Println("\nClient info:")
	fmt.Printf("Client name: %s\n", client.Name())
	fmt.Printf("Client surname: %s\n", client.Surname())
	fmt.Printf("Client account number: %d\n", client.AccountNumber())
	fmt.Printf("Client deposit money: %f\n", client.CDeposit())
	fmt.Printf("Client credit money: %f\n", client.CCredit())
}

// FindClientBySurname знаходить клієнта за прізвищем серед банків.
func FindClientBySurname(banks []*class.Bank, banksAmount int) (*class.Client, error) {
	clientName := ""
	fmt.Print("Client surname: ")
	fmt.Scanf("%s\n", &clientName)
	for i := 0; i < banksAmount; i++ {
		bank := banks[i]
		for j := 0; j < bank.ClientsAmount(); j++ {
			client := bank.Clients()[j]
			if client.Surname() == clientName {
				return client, nil
			}
		}
	}
	return new(class.Client), errors.New("client not found")
}

// FindClientById шукає клієнта за номером рахунку у списку банків.
// Параметри:
// - banks: список банків, в яких буде виконуватись пошук.
// - banksAmount: кількість банків у списку.
// Повертає: 
// - посилання на знайденого клієнта (*class.Client), якщо клієнт існує.
// - помилку (error), якщо клієнт не знайдений.
func FindClientById(banks []*class.Bank, banksAmount int) (*class.Client, error) {
	clientId := 0

	// Запит користувача для введення номера рахунку клієнта.
	fmt.Print("Client account number: ")
	fmt.Scanf("%d\n", &clientId)

	// Перебір банків.
	for i := 0; i < banksAmount; i++ {
		bank := banks[i]

		// Перебір клієнтів у банку.
		for j := 0; j < bank.ClientsAmount(); j++ {
			client := bank.Clients()[j]

			// Перевірка, чи співпадає номер рахунку.
			if client.AccountNumber() == clientId {
				fmt.Printf("Client found in bank: %s\n", bank.Name()) // Виводимо повідомлення про знайдений банк.
				return client, nil
			}
		}
	}

	// Якщо клієнта не знайдено, повертаємо помилку.
	return nil, errors.New("client not found")
}
