/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/

package main

import (
	"fmt"
	"log"
	"math"
)

func checkBottlesCount(count int) {
	if count < 0 {
		log.Fatal("Координаты должны быть >=0!")
	}
}

func checkTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	return !(x3*(y2-y1)-y3*(x2-x1) == x1*y2-x2*y1)
}

func checkTriangleSquare(x1, y1, x2, y2, x3, y3 int) float64 {
	fx1 := float64(x1)
	fy1 := float64(y1)
	fx2 := float64(x2)
	fy2 := float64(y2)
	fx3 := float64(x3)
	fy3 := float64(y3)
	//
	return math.Abs((fx2-fx1)*(fy3-fy1)-(fx3-fx1)*(fy2-fy1)) * 0.5
}

func getLength(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2)-float64(x1), 2) + math.Pow(float64(y2)-float64(y1), 2))
}

func isRightAngled(x1, y1, x2, y2, x3, y3 int) bool {
	a := getLength(x1, y1, x2, y2)
	b := getLength(x1, y1, x3, y3)
	c := getLength(x3, y3, x2, y2)

	return (math.Round(a*a+b*b) == math.Round(c*c)) || (math.Round(c*c+b*b) == math.Round(a*a)) || (math.Round(a*a+c*c) == math.Round(b*b))
}

func main() {
	var x1, y1, x2, y2, x3, y3 int
	// var x1, y1, x2, y2, x3, y3 int = 1, 1, 3, 4, 2, 2
	// var x1, y1, x2, y2, x3, y3 int = 1, 2, 4, 2, 3, 5
	// var x1, y1, x2, y2, x3, y3 int = 5, 0, 0, 0, 0, 5

	fmt.Print("Введите координаты через пробел (x1, y1, x2, y2, x3, y3): ")
	fmt.Scanf("%d %d %d %d %d %d", &x1, &y1, &x2, &y2, &x3, &y3)
	checkBottlesCount(x1)
	checkBottlesCount(y1)
	checkBottlesCount(x2)
	checkBottlesCount(y2)
	checkBottlesCount(x3)
	checkBottlesCount(y3)

	if !checkTriangle(x1, y1, x2, y2, x3, y3) {
		log.Fatal("По заданным координатам невозможно построить треугольник!")
	}

	fmt.Println("Площадь треугольника:", checkTriangleSquare(x1, y1, x2, y2, x3, y3))
	if isRightAngled(x1, y1, x2, y2, x3, y3) {
		fmt.Println("Треугольник прямоугольный")
	} else {
		fmt.Println("Треугольник не прямоугольный")
	}

	/*
		fmt.Print("Введите количество: ")
		fmt.Scanf("%d", &cn)
		checkBottlesCount(cn)
	*/
}
