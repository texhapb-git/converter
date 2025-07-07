package main

import "fmt"

func main() {
	const USD_EUR float64 = 0.9
	const USD_RUB float64 = 80.7

	const EUR_RUB float64 = USD_RUB / USD_EUR


	val, curr1, curr2 := getValues()
	calculate(val, curr1, curr2)
}

func getValues() (float64, string, string) {
	var value float64
	var inCurrency, outCurrency string

	fmt.Println("Введите сумму, исходную валюту и валюту, в которую надо перевести")
	fmt.Scan(&value, &inCurrency, &outCurrency)

	return value, inCurrency, outCurrency
} 

func calculate (value float64, inCurrency string, outCurrency string) {

}