package main

import (
	"fmt"
	"strings"
)

const RUB float64 = 1.0
const USD float64 = 80.7
const EUR float64 = 89.67

func main() {
	for {
		fmt.Println("\n=== Конвертер валют ===")
		
		// Ввод исходной валюты
		inCurrency := getCurrencyInput("Введите исходную валюту (USD/EUR/RUB): ")
		
		// Ввод суммы
		amount := getAmountInput("Введите сумму: ")
		
		// Ввод целевой валюты
		availableCurrencies := getAvailableCurrencies(inCurrency)
		prompt := fmt.Sprintf("Введите целевую валюту (%s): ", availableCurrencies)
		outCurrency := getTargetCurrencyInput(prompt, inCurrency)
		
		// Конвертация
		calculate(amount, inCurrency, outCurrency)
		
		fmt.Println("\nНажмите Enter для новой конвертации или Ctrl+C для выхода...")
		fmt.Scanln() // Ожидание Enter
	}
}

func getCurrencyInput(prompt string) string {
	var currency string
	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		
		// Валидируем валюту
		validCurrency, err := validateCurrency(currency)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		
		return validCurrency
	}
}

func getAvailableCurrencies(excludeCurrency string) string {
	allCurrencies := []string{"USD", "EUR", "RUB"}
	var available []string
	
	for _, currency := range allCurrencies {
		if currency != excludeCurrency {
			available = append(available, currency)
		}
	}
	
	return strings.Join(available, "/")
}

func getTargetCurrencyInput(prompt string, sourceCurrency string) string {
	var currency string
	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		
		// Валидируем валюту
		validCurrency, err := validateCurrency(currency)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		
		// Проверяем что целевая валюта отличается от исходной
		if validCurrency == sourceCurrency {
			fmt.Printf("Ошибка: целевая валюта не может быть такой же как исходная (%s)\n", sourceCurrency)
			continue
		}
		
		return validCurrency
	}
}

func getAmountInput(prompt string) float64 {
	var amount float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return amount
}

func getValues() (float64, string, string) {
	var value float64
	var inCurrency, outCurrency string

	fmt.Println("Введите сумму, исходную валюту и валюту, в которую надо перевести")
	fmt.Scan(&value, &inCurrency, &outCurrency)

	// Валидация исходной валюты
	validInCurrency, err := validateCurrency(inCurrency)
	if err != nil {
		fmt.Println("Ошибка исходной валюты:", err)
		return 0, "", ""
	}

	// Валидация целевой валюты
	validOutCurrency, err := validateCurrency(outCurrency)
	if err != nil {
		fmt.Println("Ошибка целевой валюты:", err)
		return 0, "", ""
	}

	return value, validInCurrency, validOutCurrency
}

func validateCurrency(currency string) (string, error) {
	upperCurrency := strings.ToUpper(currency)
	
	validCurrencies := []string{"USD", "EUR", "RUB"}
	
	for _, valid := range validCurrencies {
		if upperCurrency == valid {
			return upperCurrency, nil
		}
	}
	
	return "", fmt.Errorf("неподдерживаемая валюта: %s. Поддерживаемые валюты: USD, EUR, RUB", currency)
}

func calculate (value float64, inCurrency string, outCurrency string) {
	if inCurrency == outCurrency {
		fmt.Println(value)
		return
	}
	
	var result float64
	
	// Конвертируем в RUB (базовая валюта)
	var rubValue float64
	switch inCurrency {
	case "RUB":
		rubValue = value
	case "USD":
		rubValue = value * USD
	case "EUR":
		rubValue = value * EUR
	}
	
	// Конвертируем из RUB в целевую валюту
	switch outCurrency {
	case "RUB":
		result = rubValue
	case "USD":
		result = rubValue / USD
	case "EUR":
		result = rubValue / EUR
	}
	
	fmt.Printf("%.2f %s = %.2f %s\n", value, inCurrency, result, outCurrency)
}