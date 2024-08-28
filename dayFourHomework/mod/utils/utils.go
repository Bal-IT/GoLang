package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IsLetterOrSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && r != ' ' {
			return false
		}
	}

	//unicode.()
	return true
}

func IsDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func GetInputStr(caption string) string {
	fmt.Print(caption + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	return input
}

func GetInputInt(caption string) int {
	fmt.Print(caption + ": ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	i, err := strconv.Atoi(input)
	if err != nil {
		log.Panic(err)
	}

	return i
}
