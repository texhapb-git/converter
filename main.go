package main

import (
	"fmt"
	"strings"
)

var currencies = map[string]float64{
	"RUB": 1.0,
	"USD": 80.7,
	"EUR": 89.67,
}

func main() {
	for {
		fmt.Println("\n=== Конвертер валют ===")
		
		// Ввод исходной валюты
		inCurrency := getCurrencyInput("Введите исходную валюту (USD/EUR/RUB): ")
		
		// Ввод суммы
		amount := getAmountInput("Введите сумму: ")
		
		// Ввод целевой валюты
		availableCurrencies := getAvailableCurrencies(inCurrency, &currencies)
		prompt := fmt.Sprintf("Введите целевую валюту (%s): ", availableCurrencies)
		outCurrency := getTargetCurrencyInput(prompt, inCurrency)
		
		// Конвертация
		calculate(amount, inCurrency, outCurrency, &currencies)
		
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
		validCurrency, err := validateCurrency(currency, &currencies)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		
		return validCurrency
	}
}

func getAvailableCurrencies(excludeCurrency string, currencies *map[string]float64) string {
	var available []string
	
	for currency := range *currencies {
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
		validCurrency, err := validateCurrency(currency, &currencies)
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

func validateCurrency(currency string, currencies *map[string]float64) (string, error) {
	upperCurrency := strings.ToUpper(currency)
	
	for valid := range *currencies {
		if upperCurrency == valid {
			return upperCurrency, nil
		}
	}
	
	// Собираем список поддерживаемых валют для сообщения об ошибке
	var validCurrencies []string
	for currency := range *currencies {
		validCurrencies = append(validCurrencies, currency)
	}
	
	return "", fmt.Errorf("неподдерживаемая валюта: %s. Поддерживаемые валюты: %s", currency, strings.Join(validCurrencies, ", "))
}

func calculate (value float64, inCurrency string, outCurrency string, currencies *map[string]float64) {
	if inCurrency == outCurrency {
		fmt.Println(value)
		return
	}
	
	var result float64
	
	// Конвертируем в RUB (базовая валюта)
	rubValue := value * (*currencies)[inCurrency]
	
	// Конвертируем из RUB в целевую валюту
	result = rubValue / (*currencies)[outCurrency]
	
	fmt.Printf("%.2f %s = %.2f %s\n", value, inCurrency, result, outCurrency)
}