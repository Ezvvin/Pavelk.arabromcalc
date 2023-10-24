package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rome []string = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите пример:")
	input, _ := reader.ReadString('\n')        //вводим пример.
	input = strings.TrimSpace(input)           //Очищает все пустоты.
	text := strings.ReplaceAll(input, " ", "") //удаляем пробелы.
	if string(text[0]) == "-" {                //Проверяем, является ли первое число отрицательным.
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		panic(errors.New("Вывод ошибки, числа не могут быть отрицательными."))
	}

	switch {
	case strings.Contains(text, "+"): //Ищем знак в примере.

		arr := strings.Split(text, "+") //Превращаем пример в срез из чисел.

		typeRomans, err := checkTypeRome(arr)
		if typeRomans {
			arr[0] = romanToInteger(arr[0])
			arr[1] = romanToInteger(arr[1])
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		a1, err := strconv.Atoi(arr[0]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат первого числа неверный.", err)
			return
		}
		a2, err := strconv.Atoi(arr[1]) //Конвертируем второе число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго числа неверный.", err)
			return
		}
		if _, err := checkArguments(a1, a2); err != nil {
			fmt.Println(err)
			return
		}
		result := int(a1) + int(a2)
		if typeRomans {
			fmt.Println(integerToRoman(result))
		} else {
			fmt.Println(result)
		}
	case strings.Contains(text, "*"): //Ищем знак в примере.

		arr := strings.Split(text, "*") //Превращаем пример в срез из чисел.

		typeRomans, err := checkTypeRome(arr)
		if typeRomans {
			arr[0] = romanToInteger(arr[0])
			arr[1] = romanToInteger(arr[1])
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		a1, err := strconv.Atoi(arr[0]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат первого числа неверный.", err)
			return
		}
		a2, err := strconv.Atoi(arr[1]) //Конвертируем второе число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго числа неверный.", err)
			return
		}
		if _, err := checkArguments(a1, a2); err != nil {
			fmt.Println(err)
			return
		}
		result := int(a1) * int(a2)
		if typeRomans {
			fmt.Println(integerToRoman(result))
		} else {
			fmt.Println(result)
		}
	case strings.Contains(text, "/"): //Ищем знак в примере.

		arr := strings.Split(text, "/") //Превращаем пример в срез из чисел.

		typeRomans, err := checkTypeRome(arr)
		if typeRomans {
			arr[0] = romanToInteger(arr[0])
			arr[1] = romanToInteger(arr[1])
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		a1, err := strconv.Atoi(arr[0]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат первого числа неверный.", err)
			return
		}
		a2, err := strconv.Atoi(arr[1]) //Конвертируем второе число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго числа неверный.", err)
			return
		}
		if _, err := checkArguments(a1, a2); err != nil {
			fmt.Println(err)
			return
		}
		result := int(a1) / int(a2)
		if typeRomans {
			fmt.Println(integerToRoman(result))
		} else {
			fmt.Println(result)
		}
	case strings.Contains(text, "-"): //Ищем знак в примере.

		arr := strings.Split(text, "-") //Превращаем пример в срез из чисел.

		typeRomans, err := checkTypeRome(arr)
		if typeRomans {
			arr[0] = romanToInteger(arr[0])
			arr[1] = romanToInteger(arr[1])
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := checkLen(arr); err != nil {
			fmt.Println(err)
			return
		}
		a1, err := strconv.Atoi(arr[0]) //Конвертируем первое число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат первого числа неверный.", err)
			return
		}
		a2, err := strconv.Atoi(arr[1]) //Конвертируем второе число из string в int.
		if err != nil {
			fmt.Println("Вывод ошибки: формат второго числа неверный.", err)
			return
		}
		if _, err := checkArguments(a1, a2); err != nil {
			fmt.Println(err)
			return
		}
		result := int(a1) - int(a2)
		if typeRomans {
			if result < 1 {
				fmt.Println("Вывод ошибки: в римской системе нет отрицательных чисел.")
				return
			}
			fmt.Println(integerToRoman(result))
		} else {
			fmt.Println(result)
		}
	default:
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
	}
}

// Проверяем кол-во чисел в примере.
func checkLen(a []string) (bool, error) {
	if len(a) > 2 {
		return false, errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	return true, nil
}

// Проверяем диапазон чисел от 1 до 10 включительно.
func checkArguments(a, b int) (bool, error) { //
	if (a >= 1 && a <= 10) && (b >= 1 && b <= 10) {
		return true, nil
	}
	return false, errors.New("Вывод ошибки, числа не могут быть меньше 1 и больше 10 включительно.")
}

// Конвертируем риские числа в арабские.
func romanToInteger(s string) string {
	switch s {
	case "I":
		return "1"
	case "II":
		return "2"
	case "III":
		return "3"
	case "IV":
		return "4"
	case "V":
		return "5"
	case "VI":
		return "6"
	case "VII":
		return "7"
	case "VIII":
		return "8"
	case "IX":
		return "9"
	case "X":
		return "10"
	default:
		return "Вывод ошибки, числа не могут быть меньше I и больше X включительно."
	}
}

// Конвертируем арабские числа в римские.
func integerToRoman(res int) string {
	maxRomanNumber := 100
	if res > maxRomanNumber {
		return strconv.Itoa(res)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for res >= conversion.value { //Находим максимальное значение для результата.
			roman.WriteString(conversion.digit) //Записываем в строку римскую букву.
			res -= conversion.value             //Отнимаем это значение до тех пор, пока будет 0.
		}
	}

	return roman.String()
}

// Проверяем, являются ли два числа римскими.
func checkTypeRome(a []string) (bool, error) {
	var b, bb bool
	for i := 0; i < 10; i++ {
		if a[0] == rome[i] {
			b = true
		}
	}
	for i := 0; i < 10; i++ {
		if a[1] == rome[i] {
			bb = true
		}
	}
	if b == false && bb == false {
		return false, nil
	}
	if b == false || bb == false {
		return false, errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	return true, nil
}
