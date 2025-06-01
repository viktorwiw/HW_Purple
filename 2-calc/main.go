package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, err := scanOperation()
	if err != nil {
		fmt.Printf("Ошибка: %v.", err)
		return
	}

	digits, err := scanDigits()
	if err != nil {
		fmt.Printf("Ошибка: %v.", err)
		return
	}

	sum, err := calcSum(digits)
	if err != nil {
		fmt.Printf("Ошибка: %v.", err)
		return
	}

	switch operation {
	case "SUM":
		fmt.Printf("Сумма чисел равна: %v", sum)
	case "AVG":
		avg, err := calcAvg(sum, digits)
		if err != nil {
			fmt.Printf("Ошибка: %v.", err)
			return
		}
		fmt.Printf("Среднее из чисел равно: %v", avg)
	case "MED":
		med, err := calcMed(digits)
		if err != nil {
			fmt.Printf("Ошибка: %v.", err)
			return
		}
		fmt.Printf("Медиана равна: %v", med)
	}
}

func scanOperation() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите какую операцию хотите выполнить(AVG | SUM | MED): ")
	_ = scanner.Scan()
	operation := strings.TrimSpace(scanner.Text())

	if !isValidOperation(operation) {
		return "", errors.New("неверная операция, допустимые значения AVG, SUM, MED")
	}
	return operation, nil
}

func isValidOperation(operation string) bool {
	return operation == "AVG" || operation == "SUM" || operation == "MED"
}

func scanDigits() ([]string, error) {
	fmt.Print("Введите числа через запятую: ")
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	digits := scanner.Text()
	if strings.TrimSpace(digits) == "" {
		return nil, errors.New("вы не ввели числа")
	}
	scanDigits := strings.Split(strings.ReplaceAll(digits, " ", ""), ",")
	return scanDigits, nil
}

func calcAvg(sum float64, digits []string) (float64, error) {
	if len(digits) == 0 {
		return 0, errors.New("невозможно вычислить среднее: нет чисел")
	}
	return sum / float64(len(digits)), nil
}

func calcMed(digits []string) (float64, error) {
	if len(digits) == 0 {
		return 0, errors.New("невозможно вычислить медиану: нет чисел")
	}
	numbers := make([]float64, len(digits))
	for i, value := range digits {
		digit, err := strconv.Atoi(value)
		if err != nil {
			return 0, errors.New("невозможно преобразовать строку в числа")
		}
		numbers[i] = float64(digit)
	}

	sort.Float64s(numbers)

	lenNumbers := len(numbers)

	if lenNumbers%2 == 1 {
		return numbers[lenNumbers/2], nil
	}

	mid1 := numbers[lenNumbers/2]
	mid2 := numbers[lenNumbers/2-1]
	return (mid1 + mid2) / 2, nil
}

func calcSum(digits []string) (float64, error) {
	var sum float64
	for _, value := range digits {
		digit, err := strconv.Atoi(value)
		if err != nil {
			return 0, errors.New("невозможно преобразовать строку в числа")
		}
		sum += float64(digit)
	}
	return sum, nil
}
