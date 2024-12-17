package t4

import "testing"

// Тест для перевірки мінімального значення серед трьох чисел
func TestFindMinimum(t *testing.T) {
    result := findMinimum(1, 2, 3)  // Викликаємо функцію findMinimum
    expected := float32(1)  // Ожидаємо, що мінімальне значення буде 1
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)  // Виводимо повідомлення про помилку, якщо значення не співпадають
    }
}

// Тест для перевірки середнього значення трьох чисел
func TestCalculateAverage(t *testing.T) {
    result := calculateAverage(1, 2, 3)  // Викликаємо функцію calculateAverage
    expected := float32(2)  // Ожидаємо, що середнє значення буде 2
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)  // Виводимо повідомлення про помилку, якщо значення не співпадають
    }
}

// Тест для перевірки рівняння першого порядку: y = kx + b
func TestFirstOrderEquation(t *testing.T) {
    result := firstOrderEquation(1, 2, 3)  // Викликаємо функцію firstOrderEquation
    expected := float32(5)  // Ожидаємо, що результат рівняння буде 5 (1*2 + 3)
    if result != expected {
        t.Errorf("Тест не пройдено. Результат - %f, а вірний - %f", result, expected)  // Виводимо повідомлення про помилку, якщо значення не співпадають
    }
}
