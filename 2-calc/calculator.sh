#!/bin/bash

# Компиляция Go-кода
echo "Компиляция Go-кода..."
go build -o currency_calculator main.go

# Запуск программы
echo "Запуск калькулятора валют..."
./currency_calculator