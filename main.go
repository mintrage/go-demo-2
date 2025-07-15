package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив
// !Вывести сумму баланса в консоль

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransactions()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", balance)

}

func scanTransactions() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
