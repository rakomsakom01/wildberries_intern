package main


import (
    "fmt"
)

/*
Задание номер 8

Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/


func main() {

    // num - исходное число
    // i - номер бита, который мы хотим поменять
    // change - на что меняем
    var num uint64
    var i int
    var change bool

    fmt.Scanf("%d", &num)
    fmt.Printf("Old num:    %064b\n", num)
    fmt.Scanf("%d %t", &i, &change)

    var newNum = num
    
    if change {
        // Изменяем i - тый бит на 1

        var mask uint64

        mask = 1 << i

        newNum |= mask
    } else {
        // Изменяем i-тый бит на 0

        var mask uint64
        mask = ^ (1 << i)

        newNum &= mask
    }

    fmt.Printf("%d\n", newNum)

    fmt.Printf("Bin:\n")

    fmt.Printf("Old num:    %064b\n", num)
    fmt.Printf("New num:    %064b\n", newNum)
    fmt.Printf("Change bit: %064b\n", uint64(1 << i))

}
