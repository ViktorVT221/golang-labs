package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
	"strings"
)

func calculateTourCost(numbDay int64, country, season int, room, guide bool) float64 {
	var sum float64
	// Розрахунок вартості
	switch {
	case country == 0 && season == 0:
		sum = float64(numbDay) * 100.0
	case country == 0 && season == 1:
		sum = float64(numbDay) * 150.0
	case country == 1 && season == 0:
		sum = float64(numbDay) * 160.0
	case country == 1 && season == 1:
		sum = float64(numbDay) * 200.0
	case country == 2 && season == 0:
		sum = float64(numbDay) * 120.0
	case country == 2 && season == 1:
		sum = float64(numbDay) * 180.0
	}

	// Додавання додаткових опцій
	if guide {
		sum += (sum / 100.0) * 20.0 // 20% за індивідуального гіда
	}
	if room {
		sum += float64(numbDay) * 50.0 // Додатково за номер люкс
	}

	return sum
}

func Calculate() {
	// Створення вікна
	window := ui.NewWindow("Тур", 100, 120, false)
	window.SetMargined(false)

	// Основний контейнер
	box := ui.NewHorizontalBox()
	box.SetPadded(true)

	// Контейнер для введення даних
	info := ui.NewVerticalBox()
	info.SetPadded(true)

	// Група для введення параметрів туру
	tour := ui.NewGroup("Тур")
	inputValue := ui.NewForm()
	inputValue.SetPadded(true)

	// Поля для введення кількості днів, країни та сезону
	days := ui.NewEntry()
	countries := ui.NewCombobox()
	countries.Append("Болгарія")
	countries.Append("Німечина")
	countries.Append("Польща")
	seasons := ui.NewCombobox()
	seasons.Append("Літо")
	seasons.Append("Зима")

	// Додавання полів до форми
	inputValue.Append("Кількість днів", days, false)
	inputValue.Append("Країна", countries, false)
	inputValue.Append("Сезон", seasons, false)

	// Додавання форми до групи
	info.Append(inputValue, false)
	tour.SetChild(info)
	box.Append(tour, false)

	// Додаткова інформація (перевірки для номеру люкс і індивідуального гіда)
	addInfo := ui.NewGroup("Додаткова інформація")
	addInfo.SetMargined(true)
	checkInfo := ui.NewVerticalBox()
	checkInfo.SetPadded(true)
	deluxe := ui.NewCheckbox("Номер люкс")
	guide := ui.NewCheckbox("Індивідуальний гід")

	// Додавання чекбоксів до додаткової інформації
	checkInfo.Append(deluxe, false)
	checkInfo.Append(guide, false)
	addInfo.SetChild(checkInfo)
	info.Append(addInfo, false)

	// Розділ для результату
	result := ui.NewVerticalBox()
	result.SetPadded(true)
	res := ui.NewGroup("Сума")
	res.SetMargined(true)
	boxRes := ui.NewVerticalBox()
	boxRes.SetPadded(true)
	label := ui.NewLabel("")

	// Додавання лейблу до результату
	boxRes.Append(label, false)
	res.SetChild(boxRes)
	result.Append(res, false)

	// Панель для кнопки розрахунку
	boxButt := ui.NewHorizontalBox()
	button := ui.NewButton("Розрахувати ціну")
	boxButt.Append(button, false)
	result.Append(boxButt, false)

	// Додавання результату в основний контейнер
	box.Append(result, false)

	// Обробник події для кнопки "Розрахувати"
	button.OnClicked(func(*ui.Button) {
		countDay := strings.TrimSpace(days.Text())
		numbDay, errorD := strconv.ParseInt(countDay, 0, 64)
		errors := ui.NewWindow("ERROR", 10, 10, false)

		// Перевірка на порожнє введення
		if countDay == "" {
			ui.MsgBoxError(errors, "Помилка!", "Введіть кількість днів!")
			return
		}

		// Перевірка на правильність введених даних
		if errorD == nil {
			// Отримання вибраних параметрів
			country := countries.Selected()
			season := seasons.Selected()
			room := deluxe.Checked()
			guideSelected := guide.Checked()

			// Розрахунок суми
			sum := calculateTourCost(numbDay, country, season, room, guideSelected)

			// Виведення результату
			label.SetText(strconv.FormatFloat(sum, 'f', -1, 64) + " доларів")

		} else {
			// Повідомлення про помилку
			ui.MsgBoxError(errors, "Помилка!", "Перевірте правильність введення")
		}
	})

	// Обробник закриття вікна
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	// Встановлення основного вікна
	window.SetChild(box)

	// Показ вікна
	window.Show()
}

func main() {
	// Запуск програми
	err := ui.Main(Calculate)
	if err != nil {
		panic(err)
	}
}
