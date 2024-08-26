package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

func GetInputStr() string {
	fmt.Print("Введите код: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	return input
}

func checkInput(code string) error {
	if len(code) == 0 {
		return errors.New("количество имволов в строке должно быть больше нуля")
	}

	if len(code)%2 != 0 {
		return errors.New("количество символов в строке должно быть чётным")
	}

	return nil
}

func codeToString(code string) (string, error) {
	dicti := make(map[int]rune)
	var res string

	for idx := 0; idx <= 25; idx++ {
		dicti[idx] = rune(int('a') + idx)
	}
	dicti[26] = ' '

	for i := 0; i < len(code); i += 2 {
		// fmt.Printf("%c%c ", code[i], code[i+1])

		ch, err := strconv.Atoi(string(code[i]) + string(code[i+1]))
		if err != nil {
			return "", err
		}

		if ch > 26 {
			return "", errors.New("невозможно подобрать ключ")
		}

		res += string(dicti[ch])
	}
	return res, nil
}

func main() {
	// var code string = "220411112603141304"
	code := GetInputStr()

	err := checkInput(code)
	if err != nil {
		log.Fatal(err)
	}

	str, err := codeToString(code)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Закодированный текст:", str)
}
