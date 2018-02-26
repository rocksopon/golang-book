package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
	items map[string]int

}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
	/*if coin == "T" {
		m.insertedMoney += 10
	} else if coin == "F" {
		m.insertedMoney += 5
	} else if coin == "TW" {
		m.insertedMoney += 2
	} else if coin == "O" {
		m.insertedMoney += 1
	}*/
}

func (m *VendingMachine) SelectSD() string {
	price := m.items["SD"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "SD" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	price := m.items["CC"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "CC" + m.change(change)
}

func (m *VendingMachine) change(c int) string {
	var str string
	if c == 8 {
		return ", F, TW, O"
	}
	return str

}

func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW":2, "O":1}
	var items = map[string]int{"SD": 18, "CC": 12}
	vm := VendingMachine{coins: coins, items: items}


	fmt.Println("Inserted Money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())

	can := vm.SelectSD()
	fmt.Println(can) //SD

	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can) //CC

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can) //CC, F, TW, O
}