package main

import (
    "fmt"
    "sync"
)

/*
Задание номер 2

Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива 
(2,4,6,8,10) и выведет их квадраты в stdout.
*/

// Здесь считаются квадраты и записываются в канал
func square(amu []sync.Mutex, ch chan<- int, i int, a []int) {
    defer amu[i].Unlock()
    var arr = a[i] * a[i]
    // Ждём, когда предыдущий элемент обработается(если существует)
    if (i != 0) {
        amu[i - 1].Lock()
    }

    ch <- arr
}


func main() {
    // n -длина массива
    var n int
    fmt.Scanf("%d", &n)
    
    // a - исходный массив
    var a = make([]int, n)

    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }
    
    // Массив мютексов нужен для сохранения порядка.
    // Изначально мютексы залочены, но после вычисления очередного числа он разлочивается
    var amu = make([]sync.Mutex, n)

    // Горутины записывают полученные числа в канал, откуда основной поток их считывает
    var ch = make(chan int)

    for i := range(a) {
        amu[i].Lock()
        go square(amu, ch, i, a)
    }

    for _ = range(a) {
        fmt.Printf("%d ", <-ch)
    }
    fmt.Printf("\n")
}
