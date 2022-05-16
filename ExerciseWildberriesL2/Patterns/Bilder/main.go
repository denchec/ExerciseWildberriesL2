package main

import "ExerciseWildberriesL2/Bilder/pkg"

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")

	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}

/* Реализация паттерна "строитель" в данной программе:
Определили возможные комплектующие.
Создали базовый завод с выбранной по умолчанию комплектацией(asus) и производим компьютеры
Потом мы решили поменять комплектацию компьютеров, на комплектацию другой компании(hp) и опять производим
В самом конце мы решили обратно поменять комплектацию компьютеров на asus


1) Плюсы:
1.1) Позволяет создавать пошагово общий продукт
1.2) Позволяет использовать один и тот же код для создания различных объектов
1.3) Изолирует сложный код сборки объекта и его основной бизнес логики


2) Минусы:
2.1) Усложняет код программы из-за введения дополнительных структур, интерфейсов и т.д.
2.2) Клиент будет привязан к определенному объекту строителя т.е. он всегда будет создавать объект при помощи
данного строителя
2.3) В интерфейсе может не быть какого-то метода и по этому ему нужно будет добавить его туда вручную


3.1) Позволяет создавать сложные объекты используя "шаги" т.е. на каждом шаге производится какая-то часть общего объекта
тем самым выполняя все шаги по очереди мы формируем некий объект, который из себя представляет сложную структуру.
3.2) Строитель дает возможность использовать один и тот же код строительства объекта для получения разных представлений
этого объекта т.е.мы можем пропустить шаги или добавить дополнительные, в зависимости от необходимого объекта


4) Паттерн может использоваться на фабриках и заводах для производства похожих объектов(например компьютеров разных брендов)



*/