package main

import (
    "fmt"
)
/*
Задание номер 1

Дана структура Human (с произвольным набором полей и методов). 

Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/


// Родительская структура Human
type Human struct {
    age int
    name string
}


func (h *Human) Init(age int, name string) {
    h.age = age
    h.name = name
}


func (h *Human) Say() {
    fmt.Printf("Hello, my name is %s!\n", h.name)
}


func (h *Human) Age() {
    fmt.Printf("%d\n", h.age)
}

// Дочерняя структура Action
type Action struct {
    Human
    work string
}

func (a *Action) Do() {
    fmt.Printf("Im working in %s\n", a.work)
}


func main() {
    var action Action
    
    action.Init(10, "Vitas")
    action.work = "factory"
    // Это методы от Human
    action.Age()
    action.Say()

    // Это метод от Action
    action.Do()
}
