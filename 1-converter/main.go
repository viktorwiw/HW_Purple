package main

import "fmt"

func main() {
	const USD_EUR = 0.9
	const USD_RUB = 80
	const EUR_RUB = 91
}

func getUserValues() (float64, string, string) {
	var amount float64
	var sourceСurrency, targetСurrency string
	fmt.Print("Введите сколько хотите поменять")
	fmt.Scan(&amount)
	fmt.Print("Введите какую валюту хотите поменять (USD, EUR, RUB)")
	fmt.Scan(&sourceСurrency)
	fmt.Print("Введите какую валюту хотите получить после обмена (USD, EUR, RUB)")
	fmt.Scan(&targetСurrency)
	return amount, sourceСurrency, targetСurrency
}

func changeMoney(amount float64, sourceСurrency string, targetСurrency float64) float64 {}
