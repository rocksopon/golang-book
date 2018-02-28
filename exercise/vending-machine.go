package main

import (
	"fmt"
)

type NewVendingMachine struct {
	coins int
	items string
	coinsReturnInt int
	priceSD int
	priceCC int
}

func main() {
	vm := NewVendingMachine{}

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())

	can := vm.SelectSD()
	fmt.Println(can) //SD

	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F 
}

func (n *NewVendingMachine) InsertCoin(coins string) {
	if coins == "T" {
		n.coins += 10
	} else if coins == "F" {
		n.coins += 5
	} else if coins == "TW" {
		n.coins += 2
	} else {
		n.coins += 1
	}
}

func (n *NewVendingMachine) GetInsertedMoney() int {
	return n.coins
}

func (n *NewVendingMachine) SelectSD() string {
	n.priceSD = 18
	n.coinsReturnInt = n.coins - n.priceSD
	ReturnCoin := n.CoinReturn()
	return "SD" + ReturnCoin
}

func (n *NewVendingMachine) SelectCC() string {
	n.priceCC = 12
	n.coinsReturnInt = n.coins - n.priceCC
	ReturnCoin := n.CoinReturn()
	return "CC" + ReturnCoin
}

func (n *NewVendingMachine) CoinReturn() string {

	for i := 1; i <= (n.coinsReturnInt/10); i++ {
		fmt.Print("T ")
	}

	if n.coinsReturnInt % 10 == 5 {
		fmt.Print("F ")
	} else if n.coinsReturnInt % 10 == 2 {
		fmt.Print("TW")
	} else if n.coinsReturnInt % 10 == 1 {
		fmt.Print("O")
	}

	// Reset n.coins
	n.coins = 0
	return ""
}


