package main

import (
	"errors"
	"fmt"
)

const USD_EUR = 0.9
const USD_RUB = 80
const EUR_RUB = 91

func main() {
	for {
		sourceСurrency, amount, targetСurrency, err := getUserValues()
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Printf("Вы меняете %.0f %s на %s \n", amount, sourceСurrency, targetСurrency)
		fmt.Printf("После обмена получите %.0f %s \n", changeMoney(amount, sourceСurrency, targetСurrency), targetСurrency)
	}
}

func getUserValues() (string, float64, string, error) {
	var amount float64
	var sourceСurrency, targetСurrency string

	fmt.Print("Введите какую валюту хотите поменять (USD, EUR, RUB): ")
	fmt.Scan(&sourceСurrency)
	if sourceСurrency != "USD" && sourceСurrency != "EUR" && sourceСurrency != "RUB" {
		return "", 0, "", errors.New("Исходная Валюта неправильно задана")
	}

	fmt.Print("Введите сколько хотите поменять: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		return "", 0, "", errors.New("Количество задано неверно")
	}

	fmt.Printf("Введите какую валюту хотите получить после обмена %s: ", getValidCurrency(sourceСurrency))
	fmt.Scan(&targetСurrency)
	if !checkTargetCurrency(sourceСurrency, targetСurrency) {
		return "", 0, "", errors.New("Целевая валюта задана неверно")
	}
	return sourceСurrency, amount, targetСurrency, nil
}

func getValidCurrency(sourceСurrency string) string {
	switch sourceСurrency {
	case "USD":
		return "EUR, RUB"
	case "EUR":
		return "USD, RUB"
	case "RUB":
		return "EUR, USD"
	default:
		return ""
	}
}

func checkTargetCurrency(sourceСurrency, targetСurrency string) bool {
	switch sourceСurrency {
	case "USD":
		return targetСurrency == "EUR" || targetСurrency == "RUB"
	case "EUR":
		return targetСurrency == "USD" || targetСurrency == "RUB"
	case "RUB":
		return targetСurrency == "EUR" || targetСurrency == "USD"
	default:
		return false
	}
}

func changeMoney(amount float64, sourceСurrency string, targetСurrency string) float64 {
	switch {
	case sourceСurrency == "USD" && targetСurrency == "EUR":
		return amount * USD_EUR
	case sourceСurrency == "USD" && targetСurrency == "RUB":
		return amount * USD_RUB
	case sourceСurrency == "EUR" && targetСurrency == "USD":
		return amount / USD_EUR
	case sourceСurrency == "EUR" && targetСurrency == "RUB":
		return amount * EUR_RUB
	case sourceСurrency == "RUB" && targetСurrency == "USD":
		return amount / USD_RUB
	case sourceСurrency == "RUB" && targetСurrency == "EUR":
		return amount * EUR_RUB
	default:
		return 0
	}
}
