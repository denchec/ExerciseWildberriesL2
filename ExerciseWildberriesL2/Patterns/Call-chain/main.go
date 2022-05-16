package main

import "ExerciseWildberriesL2/Call-chain/pkg"

func main() {
	device := &pkg.Device{Name: "Device-1"}
	updateSrv := &pkg.UpdateDataService{Name: "Update-1"}
	saveSrv := &pkg.SaveDataService{}
	device.SetNext(updateSrv)
	updateSrv.SetNext(saveSrv)

	data := &pkg.Data{}
	device.Execute(data)
}

/* Реализация паттерна "цепочка вызовов" в данной программе:
Берутся произвольные данные, которые отправляются по цепочке обработки:
1) Получаем данные
2) Обновляем/изменяем данные
3) Сохраняем данные


1) Плюсы:
1.1) Уменьшает зависимость между клиентом и обработчиком т.е. каждый обработчик, независимо, выполняет свою роль и свою
внутреннюю логику.
1.2) При необходимости изменить логику обработчика - остальные не затрагиваются.
1.3) Реализует принцип единственной обязанности(каждый сервис выполняет свою роль)
1.4) А так же реализует принцип открытости и закрытости ?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!?!

2) Минусы:
2.1) Запрос может остаться не обработанным т.е. логика была нарушена и предварительная задача не была исполнена.


3) Реализация паттерна цепочка вызовов(цепочка обязанностей):
3.1) Каждое звено решает свою задачу и передает объект, после выполнения задачи, следующему звену по цепочке.

4) Паттерн используется для:
4.1) Авторизации, проверки прав пользователя и т.д. по цепочке. */
