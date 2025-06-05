#!/bin/bash

# Компиляция Go-кода
echo "Компиляция Go-кода..."
go build -o converter main.go

# Запуск программы
echo "Запуск конвертера..."
./converter