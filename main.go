package main

import (
	"bufio"   // Умеет удобно читать строки, которые ты вводишь с клавиатуры
	"fmt"     // Печать/чтение в консоль: Print, Println, Printf
	"os"      // Доступ к "системным штукам": здесь возьмём клавиатуру (os.Stdin)
	"strconv" // Преобразование текста в числа (и наоборот). Нам нужно из "27" сделать 27.
	"strings" // Работа со строками: убрать лишние пробелы/перенос строки
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите Ваше Имя: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	if name == "" {
		fmt.Println("Имя не должно быть пустым, Запустите программу еще раз")
		return
	}

	fmt.Print("Введите Ваш Возраст(полных лет): ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)

	age, err := strconv.Atoi(ageInput)
	if err != nil || age < 1 || age > 120 {
		fmt.Println("Возраст должен быть не более 120 лет")
		return
	}
	fmt.Printf("Привет, %s! Тебе %d лет.\n", name, age)

	var group string
	if age < 2 {
		group = "Младенец"
	} else if age < 12 {
		group = "Ребенок"
	} else if age < 18 {
		group = "Подросток"
	} else if age < 60 {
		group = "Взрослый"
	} else {
		group = "Пенсионер"
	}
	fmt.Println("Категория: ", group)
}
