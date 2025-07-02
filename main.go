package main

import "fmt"

func main() {
	transactions := [6]int{1, 2, 3, 4, 5, 6}
	transactionsPartial := transactions[1:5]
	transactionsNewPartial := transactionsPartial[:1]
	transactionsNewPartial[0] = 30
	transactionsNewPartial = transactionsNewPartial[0:4]
	// banks := [2]string{}

	fmt.Println(transactions)
	// banks[0] = "Тинькофф"
	// fmt.Println(banks)
	// partial := transactions[1:]
	// fmt.Println(partial)
	fmt.Println(transactionsNewPartial)
	fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))
}
