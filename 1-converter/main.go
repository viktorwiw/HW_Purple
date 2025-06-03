package main

import (
	"errors"
	"fmt"
)

const USD_EUR = 0.9
const USD_RUB = 80
const EUR_RUB = USD_RUB / USD_EUR

func main() {
	for {
		sourceСurrency, err := getUserSourceСurrency()
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}

		amount, err := getUserAmount()
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}

		targetСurrency, err := getUserTargetСurrency(sourceСurrency)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}

		fmt.Printf("Вы меняете %.0f %s на %s \n", amount, sourceСurrency, targetСurrency)

		ehchangeRates := map[string]float64{}
		calcRates(ehchangeRates)

		fmt.Printf("После обмена получите %.0f %s \n", changeMoney(amount, sourceСurrency, targetСurrency), targetСurrency)
	}
}

func getUserSourceСurrency() (string, error) {
	var sourceСurrency string
	fmt.Print("Введите какую валюту хотите поменять (USD, EUR, RUB): ")
	fmt.Scan(&sourceСurrency)
	if sourceСurrency != "USD" && sourceСurrency != "EUR" && sourceСurrency != "RUB" {
		return "", errors.New("Исходная Валюта неправильно задана")
	}
	return sourceСurrency, nil
}

func getUserAmount() (float64, error) {
	var amount float64
	fmt.Print("Введите сколько хотите поменять: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		return 0, errors.New("Количество задано неверно")
	}
	return amount, nil
}

func getUserTargetСurrency(sourceСurrency string) (string, error) {
	var targetСurrency string
	fmt.Printf("Введите какую валюту хотите получить после обмена %s: ", getValidCurrency(sourceСurrency))
	fmt.Scan(&targetСurrency)
	if !checkTargetCurrency(sourceСurrency, targetСurrency) {
		return "", errors.New("Целевая валюта задана неверно")
	}
	return targetСurrency, nil
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

func calcRates(ehchangeRates map[string]float64) {
	ehchangeRates["USD_EUR"] = USD_EUR
	ehchangeRates["USD_RUB"] = USD_RUB
	ehchangeRates["EUR_USD"] = 1 / USD_EUR
	ehchangeRates["EUR_RUB"] = EUR_RUB
	ehchangeRates["RUB_USD"] = 1 / USD_RUB
	ehchangeRates["RUB_EUR"] = 1 / EUR_RUB
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
		return amount / EUR_RUB
	default:
		return 0
	}
}
