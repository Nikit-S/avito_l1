package main

import "fmt"

type Human struct {
	Age    int
	Name   string
	Gender bool // :)
}

func (h *Human) AskAge() int { // метод
	return h.Age
}

func (h *Human) AsName() string { // метод
	return h.Name
}

//композиция
type Action struct {
	Human // так происходит «наследование» в го
}

//встраивание
type Action2 struct {
	human Human // так происходит «наследование» в го
}

func main() {
	//Создавние объекта типа Action
	A := Action{Human: Human{Age: 10, Name: "Alex", Gender: true}}
	A1 := Action2{human: Human{Age: 10, Name: "Alex", Gender: true}}
	//Вызов методов Human у типа Action
	fmt.Println(A.Name)
	fmt.Println(A.AskAge())
	fmt.Println(A1.human.Name)
	fmt.Println(A1.human.AskAge())

}
