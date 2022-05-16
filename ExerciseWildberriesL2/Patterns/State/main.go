package main

import (
	"ExerciseWildberriesL2/State/pkg"
	"fmt"
	"log"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/* Реализация паттерна состояния в данной программе:
- Мы включаем вендинговую машину и кладем в нее некоторое кол-во товара, за определенную цену
- Покупатель запрашивает информацию о наличии товара. Если он есть - вносит деньги и получает товар.
Если нет - аппарат сообщит об этом.

В каждом из состояний можно сделать лишь то, что не противоречит основному действию, например:
Если мы оплачиваем товар - то не можем в этот момент запрашивать информацию о наличии товара, не можем положить в
автомат новый товар, а так же не можем получить товар т.к. мы еще не успели оплатить


1) Плюсы:
1.1) Избавляет от множества условных операторов состояний.
1.2) Концентрирует в одном месте код, связанный с определенным состоянием.
1.3) Упрощает код контекста.

2) Минусы:
2.1) Может неоправданно усложнить код, если состояний мало и они редко меняются

3) Паттерн состояние - это паттерн, который пребывает в одном из нескольких состояний и может совершать определенные
операции. При изменении состояния - изменяются и операции, которые может производить паттерн.

4) Паттерн состояния может применяться в вендинговой машине */
