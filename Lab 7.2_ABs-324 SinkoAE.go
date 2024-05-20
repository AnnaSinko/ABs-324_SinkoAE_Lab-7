package main

import (
	"fmt"
)

// Функция для получения k-й перестановки из чисел от 1 до n
func getKthPermutation(n int, k int) string {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	return getKthPermutationHelper(numbers, k-1)
}

func getKthPermutationHelper(numbers []int, k int) string {
	n := len(numbers)
	if n == 1 {
		return fmt.Sprintf("%d", numbers[0])
	}

	// Вычисление факториала для n-1
	fact := 1
	for i := 1; i <= n-1; i++ {
		fact *= i
	}

	// Определение индекса текущего числа в перестановке
	index := k / fact

	// Получение числа по вычисленному индексу из массива numbers (это число является следующим в текущей перестановке)
	number := numbers[index]

	// Удаление использованного числа из массива
	numbers = append(numbers[:index], numbers[index+1:]...)

	// Рекурсивный вызов функции для оставшихся чисел и добавление текущего числа к результату
	return fmt.Sprintf("%d%s", number, getKthPermutationHelper(numbers, k%fact))
}

func main() {
	// Вывод сообщения для пользователя и ввод значений n и k
	var n, k int
	fmt.Print("Введите n: ")
	fmt.Scan(&n)
	fmt.Print("Введите k: ")
	fmt.Scan(&k)

	// Вывод k-й перестановки
	fmt.Printf("Выход: %s\n", getKthPermutation(n, k))
}