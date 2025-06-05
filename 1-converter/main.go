package main

import (
	"errors"
	"fmt"
)

const USD_EUR = 0.9
const USD_RUB = 80.0
const EUR_RUB = USD_RUB / USD_EUR

func main() {
	exchangeRates := map[string]float64{}
	calcRates(&exchangeRates)

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

		afterChange := changeMoney(amount, sourceСurrency, targetСurrency, &exchangeRates)
		fmt.Printf("После обмена получите %.2f %s \n", afterChange, targetСurrency)
	}
}

func getUserSourceСurrency() (string, error) {
	var sourceСurrency string
	fmt.Print("Введите какую валюту хотите поменять (USD, EUR, RUB): ")
	fmt.Scan(&sourceСurrency)
	if sourceСurrency != "USD" && sourceСurrency != "EUR" && sourceСurrency != "RUB" {
		return "", errors.New("исходная Валюта неправильно задана")
	}
	return sourceСurrency, nil
}

func getUserAmount() (float64, error) {
	var amount float64
	fmt.Print("Введите сколько хотите поменять: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		return 0, errors.New("количество задано неверно")
	}
	return amount, nil
}

func getUserTargetСurrency(sourceСurrency string) (string, error) {
	var targetСurrency string
	fmt.Printf("Введите какую валюту хотите получить после обмена %s: ", getValidCurrency(sourceСurrency))
	fmt.Scan(&targetСurrency)
	if !checkTargetCurrency(sourceСurrency, targetСurrency) {
		return "", errors.New("целевая валюта задана неверно")
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

func calcRates(exchangeRates *map[string]float64) {
	(*exchangeRates)["USD_EUR"] = USD_EUR
	(*exchangeRates)["USD_RUB"] = USD_RUB
	(*exchangeRates)["EUR_USD"] = 1 / USD_EUR
	(*exchangeRates)["EUR_RUB"] = EUR_RUB
	(*exchangeRates)["RUB_USD"] = 1 / USD_RUB
	(*exchangeRates)["RUB_EUR"] = 1 / EUR_RUB
}

func changeMoney(amount float64, sourceСurrency string, targetСurrency string, exchangeRates *map[string]float64) float64 {
	return amount * (*exchangeRates)[sourceСurrency+"_"+targetСurrency]
}
