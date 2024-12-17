
// 9 варіант (непарний)
package main

import (
	"fmt"
	"errors"
)

// Структура для товару
type Product struct {
	Name     string  // Назва товару
	Price    float64 // Ціна одиниці товару
	Cost     string  // Валюта товару
	Quantity int     // Кількість на складі
	Producer string  // Назва виробника
	Weight   float64 // Вага одиниці товару
}

// Структура для валюти
type Currency struct {
	Name    string  // Назва валюти
	ExRate  float64 // Курс валюти (кількість гривень за одиницю валюти)
}

// Конструктор для Product
// Функція для створення нового товару
func NewProduct(name, cost, producer string, price float64, quantity int, weight float64) *Product {
	return &Product{
		Name:     name,
		Price:    price,
		Cost:     cost,
		Quantity: quantity,
		Producer: producer,
		Weight:   weight,
	}
}

// Конструктор для Currency
// Функція для створення нової валюти
func NewCurrency(name string, exRate float64) *Currency {
	return &Currency{
		Name:   name,
		ExRate: exRate,
	}
}

// Set-методи для Product
// Методи для встановлення значень полів товару
func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) SetPrice(price float64) {
	p.Price = price
}

func (p *Product) SetCost(cost string) {
	p.Cost = cost
}

func (p *Product) SetQuantity(quantity int) {
	p.Quantity = quantity
}

func (p *Product) SetProducer(producer string) {
	p.Producer = producer
}

func (p *Product) SetWeight(weight float64) {
	p.Weight = weight
}

// Get-методи для Product
// Методи для отримання значень полів товару
func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetCost() string {
	return p.Cost
}

func (p *Product) GetQuantity() int {
	return p.Quantity
}

func (p *Product) GetProducer() string {
	return p.Producer
}

func (p *Product) GetWeight() float64 {
	return p.Weight
}

// Метод для перерахунку ціни в гривнях за курсом валюти
func (p *Product) GetPriceIn(currency *Currency) float64 {
	return p.Price * currency.ExRate
}

// Метод для обчислення загальної вартості всіх товарів на складі
func (p *Product) GetTotalPrice(currency *Currency) float64 {
	return float64(p.Quantity) * p.GetPriceIn(currency)
}

// Метод для обчислення загальної ваги всіх товарів на складі
func (p *Product) GetTotalWeight() float64 {
	return float64(p.Quantity) * p.Weight
}

// Функція для зчитування масиву продуктів
// Зчитує з клавіатури інформацію про n товарів
func ReadProductsArray(n int) ([]Product, error) {
	products := make([]Product, n)
	for i := 0; i < n; i++ {
		var name, cost, producer string
		var price, weight float64
		var quantity int
		fmt.Printf("Введіть дані для товару %d:\n", i+1)
		fmt.Print("Назва товару: ")
		if _, err := fmt.Scan(&name); err != nil {
			return nil, err
		}
		fmt.Print("Вартість одиниці товару: ")
		if _, err := fmt.Scan(&price); err != nil {
			return nil, err
		}
		fmt.Print("Валюта (грн або інша): ")
		if _, err := fmt.Scan(&cost); err != nil {
			return nil, err
		}
		fmt.Print("Кількість на складі: ")
		if _, err := fmt.Scan(&quantity); err != nil {
			return nil, err
		}
		fmt.Print("Назва виробника: ")
		if _, err := fmt.Scan(&producer); err != nil {
			return nil, err
		}
		fmt.Print("Вага одиниці товару: ")
		if _, err := fmt.Scan(&weight); err != nil {
			return nil, err
		}

		// Додаємо продукт у масив
		products[i] = *NewProduct(name, cost, producer, price, quantity, weight)
	}
	return products, nil
}

// Функція для виведення одного товару
// Виводить на екран інформацію про товар
func PrintProduct(p Product) {
	fmt.Printf("Назва товару: %s\n", p.GetName())
	fmt.Printf("Ціна: %.2f %s\n", p.GetPrice(), p.GetCost())
	fmt.Printf("Кількість: %d\n", p.GetQuantity())
	fmt.Printf("Виробник: %s\n", p.GetProducer())
	fmt.Printf("Вага: %.2f кг\n", p.GetWeight())
}

// Функція для виведення всіх товарів
// Виводить на екран інформацію про всі товари
func PrintProducts(products []Product) {
	for _, p := range products {
		PrintProduct(p)
		fmt.Println("-------------")
	}
}

// Функція для пошуку найдешевшого та найдорожчого товару
// Повертає найдешевший та найдорожчий товар
func GetProductsInfo(products []Product) (Product, Product, error) {
	if len(products) == 0 {
		return Product{}, Product{}, errors.New("немає товарів")
	}

	cheapest := products[0]
	dearest := products[0]

	// Перевіряємо кожен товар на найдешевший та найдорожчий
	for _, p := range products[1:] {
		if p.GetPrice() < cheapest.GetPrice() {
			cheapest = p
		}
		if p.GetPrice() > dearest.GetPrice() {
			dearest = p
		}
	}

	return cheapest, dearest, nil
}

func main() {
	// 9 непарний варіант
	// Прочитати кількість товарів
	var n int
	fmt.Print("Введіть кількість товарів: ")
	fmt.Scan(&n)

	// Прочитати масив продуктів
	products, err := ReadProductsArray(n)
	if err != nil {
		fmt.Println("Помилка при введенні даних:", err)
		return
	}

	// Вивести всі продукти
	PrintProducts(products)

	// Знайти найдешевший та найдорожчий товар
	cheapest, dearest, err := GetProductsInfo(products)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	// Вивести результати
	fmt.Printf("\nНайдешевший товар:\n")
	PrintProduct(cheapest)

	fmt.Printf("\nНайдорожчий товар:\n")
	PrintProduct(dearest)
}
