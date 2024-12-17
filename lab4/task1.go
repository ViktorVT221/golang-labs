package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
	"strings"
)

func calculateCost(height, width int64, material, glass int, withSill bool) float64 {
	var sum float64
	// Розрахунок вартості в залежності від вибраних параметрів
	switch {
	case material == 0 && glass == 0:
		sum = float64(height*width) * 2.5
	case material == 1 && glass == 0:
		sum = float64(height*width) * 0.5
	case material == 2 && glass == 0:
		sum = float64(height*width) * 1.5
	case material == 0 && glass == 1:
		sum = float64(height*width) * 3.0
	case material == 1 && glass == 1:
		sum = float64(height*width) * 1.0
	case material == 2 && glass == 1:
		sum = float64(height*width) * 2.0
	}

	if withSill {
		sum += 350
	}
	return sum
}

func Count() {
	// Створення вікна
	window := ui.NewWindow("Калькулятор", 300, 250, false)
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	// Панель введення параметрів
	parameters := ui.NewHorizontalBox()
	parameters.SetPadded(true)

	// Група для введення розміру вікна
	input := ui.NewGroup("Розмір вікна")
	input.SetMargined(true)

	// Група для введення даних
	inputBox := ui.NewHorizontalBox()
	inputBox.SetPadded(true)

	// Форма для введення розмірів та матеріалу
	inputValue := ui.NewForm()
	inputValue.SetPadded(true)

	// Поля для введення ширини та висоти
	h := ui.NewEntry()
	w := ui.NewEntry()

	// Комбо-бокс для вибору матеріалу
	mater := ui.NewCombobox()
	mater.Append("Дерево")
	mater.Append("Метал")
	mater.Append("Металопластик")

	// Додавання полів у форму
	inputValue.Append("Ширина, см", h, false)
	inputValue.Append("Висота, см", w, false)
	inputValue.Append("Матеріали", mater, false)

	// Вкладення форми в групу
	inputBox.Append(inputValue, false)
	input.SetChild(inputBox)

	// Група для вибору типу склопакету
	glass := ui.NewGroup("Cклопакет")
	glass.SetMargined(true)

	// Вибір типу склопакету
	boxGlass := ui.NewVerticalBox()
	boxGlass.SetPadded(true)

	glassValue := ui.NewCombobox()
	glassValue.Append("Однокамерний")
	glassValue.Append("Двокамерний")

	// Перевірка наявності підвіконня
	check := ui.NewCheckbox("Підвіконня")

	// Додавання елементів до групи
	boxGlass.Append(glassValue, false)
	boxGlass.Append(check, false)
	glass.SetChild(boxGlass)

	// Додавання груп до панелі параметрів
	parameters.Append(input, false)
	parameters.Append(glass, false)

	// Додавання панелі параметрів в головний бокс
	box.Append(parameters, false)

	// Панель для результатів
	result := ui.NewHorizontalBox()
	result.SetPadded(true)

	// Панель для текстового результату
	boxText := ui.NewHorizontalBox()
	text := ui.NewLabel("")
	boxText.Append(text, false)
	result.Append(boxText, false)

	// Панель для кнопок
	boxButt := ui.NewVerticalBox()
	boxButt.SetPadded(true)

	// Кнопка для розрахунку
	button := ui.NewButton("Розрахувати")
	boxButt.Append(button, false)

	// Вікно для помилок
	errors := ui.NewWindow("ERROR", 20, 20, false)

	// Обробник події для кнопки
	button.OnClicked(func(*ui.Button) {
		// Отримання вибору користувача
		choosen_mat := mater.Selected()
		choosen_gl := glassValue.Selected()

		// Отримання значень висоти і ширини
		height := strings.TrimSpace(h.Text())
		width := strings.TrimSpace(w.Text())

		// Перевірка на порожні значення
		if height == "" || width == "" {
			ui.MsgBoxError(errors, "Помилка!", "Всі поля мають бути заповнені!")
			return
		}

		// Перетворення значень на числа
		costH, errorH := strconv.ParseInt(height, 0, 64)
		costW, errorW := strconv.ParseInt(width, 0, 64)

		// Перевірка на правильність введених значень
		if errorH == nil && errorW == nil {
			// Розрахунок вартості
			sum := calculateCost(costH, costW, choosen_mat, choosen_gl, check.Checked())

			// Виведення результату
			text.SetText(strconv.FormatFloat(sum, 'f', -1, 64) + " грн")
		} else {
			// Виведення помилки при невірних даних
			ui.MsgBoxError(errors, "Помилка!", "Перевірте введені дані")
		}
	})

	// Додавання кнопки до результатів
	result.Append(boxButt, true)

	// Додавання результатів в головний бокс
	box.Append(result, true)

	// Встановлення головного вікна
	window.SetChild(box)

	// Показ вікна
	window.Show()
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
}

func main() {
	err := ui.Main(Count)
	if err != nil {
		panic(err)
	}
}
