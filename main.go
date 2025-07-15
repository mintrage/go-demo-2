package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив

func main() {
	tr1 := []int{1, 2, 3}
	tr2 := []int{4, 5, 6}
	tr1 = append(tr1, tr2...)
	fmt.Print(tr1)
	// transactions := []int{0, 20, 35}
	// temp := transactions
	// transactions = append(transactions, 100)
	// fmt.Println(temp)
	// fmt.Println(transactions)
	var transactions []float64
	var input float64
	for {
		fmt.Println("Введите транзакцию:")
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		transactions = append(transactions, input)

	}
	fmt.Println(transactions)
}
