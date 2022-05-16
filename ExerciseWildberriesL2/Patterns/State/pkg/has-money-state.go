package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (hasMoney *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (hasMoney *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (hasMoney *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (hasMoney *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	hasMoney.vendingMachine.itemCount = hasMoney.vendingMachine.itemCount - 1
	if hasMoney.vendingMachine.itemCount == 0 {
		hasMoney.vendingMachine.setState(hasMoney.vendingMachine.noItem)
	} else {
		hasMoney.vendingMachine.setState(hasMoney.vendingMachine.hasItem)
	}
	return nil
}
