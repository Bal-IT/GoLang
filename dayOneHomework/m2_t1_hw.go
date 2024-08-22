package main

import "fmt"

/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

func main() {
	var d1, d2, d3 int
	var n1, n2, n3 int
	fmt.Println("Введите 3 числа:")
	fmt.Scanf("%d\n", &d1)
	fmt.Scanf("%d\n", &d2)
	fmt.Scanf("%d\n", &d3)

	if d1 > d2 {
		if d1 > d3 {
			n3 = d1 // максимальное
			if d2 > d3 {
				n2 = d2
				n1 = d3
			} else {
				n2 = d3
				n1 = d2
			}
		} else {
			n3 = d3 // максимальное
			if d2 > d1 {
				n2 = d2
				n1 = d1
			} else {
				n2 = d1
				n1 = d2
			}
		}
	} else {
		if d2 > d3 {
			n3 = d2 // максимальное
			if d3 > d1 {
				n2 = d3
				n1 = d1
			} else {
				n2 = d1
				n1 = d3
			}
		} else {
			n3 = d3 // максимальное
			if d2 > d1 {
				n2 = d2
				n1 = d1
			} else {
				n2 = d1
				n1 = d2
			}
		}
	}

	fmt.Printf("%d, %d, %d\n", n1, n2, n3)
}
