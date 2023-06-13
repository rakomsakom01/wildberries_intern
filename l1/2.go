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


func square(wg *sync.WaitGroup, i int, a []int) {
    defer wg.Done()
    a[i] = a[i] * a[i]
}


func main() {
    // n - длина масива
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

    for i := range(a) {
        fmt.Printf("%d ", a[i])
    }
    fmt.Printf("\n")
}
