package main

import (
	"fmt"
	"sort"
)

// Функция для вычисления размера предмета как произведение длины и ширины (так как предметы прямоугольные)
func calculateItemSize(length, width int) int {
	return length * width
}

// Рекурсивный алгоритм, который будет проверять все возможные комбинации предметов, которые можно использовать для перевеса веревки.
func canBalance(ropeLength int, items [][]int, index, leftSum, rightSum int) bool {
	// Базовый случай: если веревка полностью занята (после каждого перемещения предмета)
	if leftSum == rightSum || leftSum+rightSum == 2*ropeLength {
		return true // если все условия выполняются (предметы могут быть равномерно распределены по обеим веревкам веревки)
	}

	// Базовый случай: если веревка не занята и предметы закончились
	if index == len(items) {
		return false // функция возвращает false
	}

	// Попытка перевесить предмет на левой веревке
	if canBalance(ropeLength, items, index+1, leftSum+calculateItemSize(items[index][0], items[index][1]), rightSum) {
		return true // веревка может быть перевешена и предметы могут быть равномерно распределены по обеим веревкам
	}

	// Попытка перевесить предмет на правой веревке
	if canBalance(ropeLength, items, index+1, leftSum, rightSum+calculateItemSize(items[index][0], items[index][1])) {
		return true // веревка может быть перевешена и предметы могут быть равномерно распределены по обеим веревкам
	}

	// Если предмет нельзя перевесить ни на одной веревке и равномерное распределение не удается
	return canBalance(ropeLength, items, index+1, leftSum, rightSum) // пытается перевесить предмет на одной из частей веревки, не влияя на сумму размеров предметов
}

func main() {
	ropeLength := 10 // задается длина веревки
	items := [][]int{{2, 2}, {3, 2}, {2, 3}, {2, 2}, {2, 3}} // предметы заданы как длина и ширина и их количество(количество элементов в векторе)

	// Сортируем предметы по убыванию размера, чтобы начать с самых больших
	sort.Slice(items, func(i, j int) bool {
		return calculateItemSize(items[i][0], items[i][1]) > calculateItemSize(items[j][0], items[j][1])
	})

	// Добавляем украденный предмет
	stolenItem := []int{3, 2} // предмет украденный
	items = append(items, stolenItem)

	// Проверяем, можно ли перевесить веревку и выводим результат
	if canBalance(ropeLength, items, 0, 0, 0) {
		fmt.Println("Да, веревку можно уравновесить.")
	} else {
		fmt.Println("Нет, веревку нельзя уравновесить.")
	}
}