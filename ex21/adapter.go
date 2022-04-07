package main

import "fmt"

type VendingMachine interface {
	insertDollar()
}

type AmericaVendMach struct{}

func (vm *AmericaVendMach) insertDollar() {
	fmt.Println("you've got 1 'snickers'")
}

type RussianVendMach struct{}

func (vm *RussianVendMach) insertRuble() {
	fmt.Println("you've got 1 batonchik")
}

type DollarRubAdapter struct {
	*RussianVendMach
}

func (vm *DollarRubAdapter) insertDollar() {
	vm.insertRuble()
	vm.insertRuble()
	vm.insertRuble()
}

type client struct{}

func (c *client) insertCoin(vm VendingMachine) {
	vm.insertDollar()
}

/*
ну тут кажется все объяснимо по названиям
существуют вендинговые аппараты
в частности на рынке был американский, который принимал доллары
клиент умеет делать одну задачу — вставлять монетку

через какое-то время появляется Отечественный вендинговый аппарат
в него тоже можно вставлять монетку
но только в первом случае мы можем вставить доллар а в отечесвенный соотвественно рубль
что делаьб

ме делаем адаптер, у которого есть метод приема доллара с последующим запуском
отечетсвенного аппарата
*/

func main() {
	americanClient := &client{}
	avm := &AmericaVendMach{}
	rvm := &RussianVendMach{}
	adapter := &DollarRubAdapter{RussianVendMach: rvm}

	americanClient.insertCoin(avm)
	americanClient.insertCoin(adapter)
}
