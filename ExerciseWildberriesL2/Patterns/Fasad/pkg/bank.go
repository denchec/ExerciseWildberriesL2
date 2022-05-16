package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Println("[Банк] Запрос остатка по карте", cardNumber)
	time.Sleep(time.Millisecond * 800)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}

		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств на счету")
		}
	}
	fmt.Println("[Банк] Остаток положительный")
	return nil
}
