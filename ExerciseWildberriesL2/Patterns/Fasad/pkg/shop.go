package pkg

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name    string
	Product []Product
}

func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 700)

	err := user.Card.CheckBalance()

	if err != nil {
		return err
	}

	fmt.Printf("[Магазин] Проверка - может ли пользователь %s купить товар \n", user.Name)
	time.Sleep(time.Millisecond * 700)

	for _, prod := range shop.Product {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] У пользователя недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар %s - куплен! \n", prod.Name)
	}
	return nil
}
