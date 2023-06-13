package main

import (
    "fmt"
)

/*
Задание номер 3

Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func square(ch chan<- int, i int, a []int) {
    var arr = a[i] * a[i]

    ch <- arr
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

    // Горутины записывают полученные числа в канал, откуда основной поток их считывает и суммирует
    // В этом случае, кстати, нам порядок не важен
    var ch = make(chan int)

    for i := range(a) {
        go square(ch, i, a)
    }

    var sum int
    for _ = range(a) {
        sum += <-ch
    }
    fmt.Printf("%d\n", sum)
}
