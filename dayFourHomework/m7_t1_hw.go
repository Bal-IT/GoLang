package main

import (
	"day4/mod/orders"
	"day4/mod/utils"
	"fmt"
)

/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/

var errA []error

func addError(err error) {
	if err != nil {
		errA = append(errA, err)
	}
}

func main() {
	var order orders.Order
	/*
		order.ProductName = "Компьютер"
		order.Count = 1
		order.FullName = "Арутюнов "
		order.Phone = "1123456789"

		order.Index = "143909"
		order.City = "Балашиха"
		order.Street = "Фучика"
		order.Home = "25"
		order.Apartment = ""
	*/

	order.ProductName = utils.GetInputStr("Введите название продукта")
	order.Count = utils.GetInputInt("Введите количество")
	order.FullName = utils.GetInputStr("Введите ФИО")
	order.Phone = utils.GetInputStr("Введите телефон")
	order.Index = utils.GetInputStr("Введите индекс")
	order.City = utils.GetInputStr("Введите город")
	order.Street = utils.GetInputStr("Введите улицу")
	order.Home = utils.GetInputStr("Введите дом")
	order.Apartment = utils.GetInputStr("Введите квартиру")

	addError(order.CheckName())
	addError(order.CheckCount())
	addError(order.CheckFullName())
	addError(order.CheckPhone())

	addError(order.CheckIndex())
	addError(order.CheckCity())
	addError(order.CheckStreet())
	addError(order.CheckHome())
	addError(order.CheckApartment())

	if len(errA) == 0 {
		order.PrintInvoice()
	} else {
		fmt.Println("Ошибки ввода:")
		for _, e := range errA {
			fmt.Println(e)
		}
	}
}
