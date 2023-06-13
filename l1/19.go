package main

import (
    "fmt"
)

/*
Задание номер 19

Разработать программу, которая переворачивает подаваемую на ход строку 
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

// ревёрс
func reverse(str string) string {
    var raw = []rune(str)
    // Преобразуем строку, чтобы не было проблем как в 15 задании
    var rawReverse = make([]rune, len(raw))

    for i := range(raw) {
        rawReverse[len(raw) - 1 - i] = raw[i]
    }

    return string(rawReverse)
}


func main() {
    var s string

    fmt.Scanf("%s", &s)

    fmt.Printf("%s\n", reverse(s))

}
