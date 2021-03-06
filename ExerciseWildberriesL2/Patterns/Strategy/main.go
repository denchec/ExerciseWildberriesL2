package main

import "ExerciseWildberriesL2/Strategy/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	nav := pkg.Navigator{}
	for _, strategy := range strategies {
		nav.Setstrategy(strategy)
		nav.Route(start, end)
	}
}

/* Реализация паттерна "стратегия" в данной программе:
Программа является Навигатором т.е.:
В ней производится расчет времени, которое необходимо будет потратить при передвижении из одной точки в другую.
Так же учитываются разные факторы(например: пробки)
Полученная в ходе расчетов информация выводится на экран, а пользователь уже сам решает, какой путь для него лучше

1) Плюсы:
1.1) Возможность замены алгоритмов на лету.
1.2) Уход от наследования и делегирования какому-то непосредственному алгоритму.
1.3) Реализует принцип открытости и закрытости ?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!

2) Минусы:
2.1) Усложнение программы за счет дополнительных объектов
2.2) Клиент должен знать в чем состоит разница в стратегиях, чтобы выбрать подходящую

3) Применимость:
3.1) Когда нужно использовать разные варианты алгоритма внутри одного объекта - Навигатор.
3.2) Когда есть множество похожих объектов отличающихся определенным поведением. */
