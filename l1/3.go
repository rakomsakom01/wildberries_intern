package main

import (
    "fmt"
    "sync"
)

/*
Задание номер 3

Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func square(wg *sync.WaitGroup, i int, a []int) {
    defer wg.Done()
    a[i] = a[i] * a[i]
}


func main() {
    // n - длина массива
    var n int
    fmt.Scanf("%d", &n)
    // a - исходный массив
    var a = make([]int, n)

    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }

    // Wait group нужен для того, чтобы процесс вывода чисел начался после окончания работ горутин
    var wg sync.WaitGroup

    for i := range(a) {
        wg.Add(1)
        go square(&wg, i, a)
    }
    
    wg.Wait()

    var sum int

    for i := range(a) {
        sum += a[i]
    }
    fmt.Printf("%d\n", sum)
}
