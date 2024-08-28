package orders

import (
	"day4/mod/utils"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Order struct {
	ProductName string
	Count       int
	FullName    string
	Phone       string
	Address
}

type Address struct {
	Index     string
	City      string
	Street    string
	Home      string
	Apartment string
}

func (o Order) CheckName() error {
	if utf8.RuneCountInString(o.ProductName) < 1 || utf8.RuneCountInString(o.ProductName) > 100 {
		return errors.New("наименование товара должно содержать от 1 до 100 символов")
	}
	return nil
}

func (o Order) CheckCount() error {
	if o.Count < 1 {
		return errors.New("количество должно быть больше 0")
	}
	return nil
}

func (o Order) CheckFullName() error {
	if utf8.RuneCountInString(o.FullName) < 1 || !utils.IsLetterOrSpace(o.FullName) {
		return errors.New("ФИО покупателя должно содержать только буквы")
	}
	return nil
}

func (o Order) CheckPhone() error {
	if utf8.RuneCountInString(o.Phone) != 10 || !utils.IsDigit(o.Phone) {
		return errors.New("контактный телефон должен содержать 10 цифр")
	}

	if o.Phone[0] == '0' {
		return errors.New("телефон не может начинаться с 0")
	}
	return nil
}

func (o Order) PrintInvoice() {
	fmt.Println("Название продукта: ", o.ProductName)
	fmt.Println("Количество: ", o.Count)
	fmt.Println("ФИО: ", o.FullName)
	fmt.Println("Телефон: ", o.Phone)
	fmt.Println("Индекс: ", o.Index)
	fmt.Println("Город: ", o.City)
	fmt.Println("Улица: ", o.Street)
	fmt.Println("Дом: ", o.City)
	fmt.Println("Квартира: ", o.Apartment)
}

func (a Address) CheckIndex() error {
	if utf8.RuneCountInString(a.Index) != 6 || !utils.IsDigit(a.Index) {
		return errors.New("индекс должен содержать 6 цифр")
	}
	return nil
}

func (a Address) CheckCity() error {
	if a.City == "" {
		return errors.New("город должен быть заполнен")
	}
	return nil
}

func (a Address) CheckStreet() error {
	if a.Street == "" {
		return errors.New("улица должна быть заполнена")
	}
	return nil
}

func (a Address) CheckHome() error {
	if a.Home == "" {
		return errors.New("дом должен быть заполнен")
	}
	return nil
}

func (a Address) CheckApartment() error {
	if a.Apartment == "" {
		return errors.New("кваротира должна быть заполнена")
	}
	return nil
}
