package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (hasItem *HasItemState) RequestItem() error {
	if hasItem.vendingMachine.itemCount == 0 {
		hasItem.vendingMachine.setState(hasItem.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd \n")
	hasItem.vendingMachine.setState(hasItem.vendingMachine.itemRequested)
	return nil
}

func (hasItem *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items edded \n", count)
	hasItem.vendingMachine.IncrementItemCount(count)
	return nil
}

func (hasItem *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (hasItem *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
