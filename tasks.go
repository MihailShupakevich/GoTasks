package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(_2task("AAADDDGFFFAAR"))
	//fmt.Println(getMatchedNumbers([]int{1, 3, 3, 5}, []int{3, 5, 5, 6}))
	//fmt.Println(createBoard(8))
	fmt.Println(cleaned([]string{"hello", "bcda", "bruh"}))
}
func _2task(x string) string {
	// 2
	// Реализовать алгоритм сжатия строки "AAADDDGFFFAAR" → "3A3DG3F2AR"
	newStringWords := ""
	repeaterWord := 1
	for i := 0; i < len(x); i++ {
		if i < len(x)-1 && x[i] == x[i+1] {
			repeaterWord++
			continue

		}
		if repeaterWord > 1 {
			newStringWords += fmt.Sprintf("%d", repeaterWord) + string(x[i])
			repeaterWord = 1
		} else {
			newStringWords += string(x[i])
		}

	}

	return newStringWords
}

type Arrays struct {
	matched   []int
	different []int
}

func getMatchedNumbers(matched, different []int) Arrays {

	//	Реализовать функцию getMatchedNumbers. На вход поступают 2 отсортированных массива чисел Необходимо вернуть struct с двумя массивам:
	//		matched - массив цифр, в котором есть одинаковые элементы из
	//		двух входных массивов
	//		Учесть!!! Цифры не должны дублироваться.
	//		different - массив цифр, которые находятся в одном из массивов,
	//		но в другом его не было
	//		пример:
	//		на вход [1,3,3,5]; [3,5,5,6]
	//		на выходе {matched: [3,5], different: [1,6]}

	var result Arrays
	var matchedMap = make(map[int]bool)
	var differentMap = make(map[int]bool)

	for _, v := range matched {
		matchedMap[v] = true
	}
	for _, v := range different {
		differentMap[v] = true
	} //здесь 2 массива отсортированы на повторение

	for k, _ := range matchedMap {
		if differentMap[k] {
			result.matched = append(result.matched, k)
		} else {
			result.different = append(result.different, k)
		}
	}
	for k, _ := range differentMap {
		if !matchedMap[k] {
			result.different = append(result.different, k)
		}
	}
	return result
}

func createBoard(countOfBox int) string {
	initialValueOfBox := countOfBox
	result := ""
	countOfBox = countOfBox * countOfBox
	for i := 1; i <= countOfBox; i++ {
		if i%2 != 0 {
			result += "#"
		} else {
			result += "_"
		}
		if i%initialValueOfBox == 0 {
			result += "\n"
		}

	}
	return result
}

func cleaned(x5 []string) []string {

	for i := range x5 {
		r := []rune(x5[i])
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

	}
	sort.Strings(x5)
	return x5
}

// Анаграммы – это слова, у которых те же буквы в том же количестве,
// но они располагаются в другом порядке.
// Например:
// nap - pan
// ear - are - era
// cheaters - hectares - teachers
// Напишите функцию clean(arr), которая возвращает массив слов,
// очищенный от анаграмм.
// Из каждой группы анаграмм должно остаться только одно слово,
// не важно какое.
