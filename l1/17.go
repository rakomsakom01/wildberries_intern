package main

import (
    "fmt"
)

/*
Задание номер 17

Реализовать бинарный поиск встроенными методами языка.
*/

// Бинарный поиск
func binSearch(n int, a []int) int {
    var left, right int

    left = 0
    right = len(a)

    for right > left + 1 {
        var middle = (left + right) / 2

        if a[middle] > n {
            right = middle
        } else {
            left = middle
        }
    }

    return left
}


func main() {
    var n, val int

    fmt.Scanf("%d %d", &n, &val)

    var a = make([]int, n)

    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }

    index := binSearch(val, a)

    fmt.Printf("%d %d\n", index, a[index])
}
