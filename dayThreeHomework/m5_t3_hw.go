package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

const (
	digits    string = "0123456789"
	lowercase string = "abcdefghiklmnopqrstvxyz"
	uppercase string = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special   string = "_!@#$%^&"
)

func GetInputStr() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	return input
}

func charInSet(char rune, charSet string) bool {
	for _, ch := range charSet {
		if char == ch {
			return true
		}
	}
	return false
}

func checkInvalidChars(pass string) error {
	for _, ch := range pass {
		if !charInSet(ch, digits+lowercase+uppercase+special) {
			return errors.New("присутствует недопустимый символ")
		}
	}
	return nil
}

func checkChars(pass string, charSet string) bool {
	for _, ch := range pass {
		if charInSet(ch, charSet) {
			return true
		}
	}
	return false
}

func checkPass(pass string) error {
	if len(pass) < 8 || len(pass) > 15 {
		return errors.New("пароль должен содержать от 8 до 15 символов")
	}

	err := checkInvalidChars(pass)
	if err != nil {
		return err
	}

	if !checkChars(pass, digits) {
		return errors.New("пароль должен содержать цыфры")
	}

	if !checkChars(pass, lowercase) {
		return errors.New("пароль должен содержать латинские буквы в нижнем регистре")
	}

	if !checkChars(pass, uppercase) {
		return errors.New("пароль должен содержать латинские буквы в верхнем регистре")
	}

	if !checkChars(pass, special) {
		return errors.New("пароль должен содержать спецсимволы из набора _!@#$%^&")
	}

	return nil
}

func main() {
	var err error

	for tr := 1; tr <= 5; tr++ {
		fmt.Println("Поавтка", tr)
		fmt.Print("Введите пароль: ")

		pass := GetInputStr()
		err = checkPass(pass)

		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
		fmt.Println()
	}

	if err == nil {
		fmt.Println("Пароль принят!")
	} else {
		fmt.Println("Пароль не принят!")
	}

}
