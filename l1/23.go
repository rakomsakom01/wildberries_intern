package main

import (
    "fmt"
)

/*
Задание номер 23

Удалить i-ый элемент из слайса.
*/

func erase(a []int, deleteIndex int) []int {
    if (deleteIndex < 0 || deleteIndex >= len(a)) {
        return a
    }

    // Сдвигаем элементы на мето удалённого
    for i := deleteIndex; i < len(a) - 1; i++ {
        a[i] = a[i + 1]
    }
    // обрезаем срез
    a = a[:len(a) - 1]

    return a
}


func main() {
    var n, deleteIndex int

    fmt.Scanf("%d %d", &n, &deleteIndex)

    var a = make([]int, n)

    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }

    a = erase(a, deleteIndex)

    for i := range(a) {
        fmt.Printf("%d ", a[i])
    }

    fmt.Printf("\n")

}
