package main

import (
    "fmt"
    "math/rand"
)

/*
Задание номер 14

Разработать программу, которая в рантайме способна определить тип переменной: 
int, string, bool, channel из переменной типа interface{}.
*/


func defineType(interf interface{}) {
    switch interf.(type) {
    case int:
        fmt.Printf("it's int!\n")
    case string:
        fmt.Printf("it's string!\n")
    case bool:
        fmt.Printf("it's bool!\n")
    case chan int:
        fmt.Printf("it's chan int!\n")
    case chan string:
        fmt.Printf("it's chan string!\n")
    case chan bool:
        fmt.Printf("it's chan bool!\n")
    }
}


func main() {
    // n -количество тестов
    var n int

    fmt.Scanf("%d", &n)

    for i := 0; i < n; i++ {
        var interf interface{}
        
        // Генерируем данные
        r := rand.Intn(6)

        if r == 0 {
            interf = 23

            fmt.Printf("generate int\n")
        } else if r == 1 {
            interf = "ab"

            fmt.Printf("generate string\n")
        } else if r == 2 {
            interf = true

            fmt.Printf("generate bool\n")
        } else if r == 3 {
            interf = make(chan int)

            fmt.Printf("generate chan int\n")
        } else if r == 4 {
            interf = make(chan string)

            fmt.Printf("generate chan string\n")
        } else if r == 5 {
            interf = make(chan bool)

            fmt.Printf("generate chan bool\n")
        }

        // затем определяем тип переменной
        defineType(interf)
    }
}
