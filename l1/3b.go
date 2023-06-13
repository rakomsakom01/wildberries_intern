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

func reduction(wg *sync.WaitGroup, windowSize int, i int, a []int) {
    var n = len(a)

    defer wg.Done()

    if i + windowSize < n {
        a[i] = a[i] + a[i + windowSize]
    }
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

    var wg sync.WaitGroup

    // Здесь мы получаем массив квадратов
    wg.Add(len(a))
    for i := range(a) {
        go square(&wg, i, a)
    }
    
    wg.Wait()

    var windowSize = 1

    for windowSize < n {
        windowSize <<= 1
    }
    
    windowSize >>= 1
    // windowSize = 2^n, где n - такое, что n максимальное и windowSize < длина массива
    
    
    // Здесь мы суммируем числа с помощью parallel reduction 
    // для начала мы суммируем пары(если есть) чисел a[i] и a[i + windowSize] а результат
    // записываем в a[i]. Итого сумма первых windowSize чисел массива a равно сумме исходного массива
    // Затем мы уменьшаем windowSize в два раза и ещё раз суммируем. И так повторяем до тех пор, 
    // пока в windowSize не станет равным нулю. Ответ будет содержаться в a[0]
    for windowSize != 0 {
        wg.Add(windowSize)
        for i := 0; i < windowSize; i++ {
            go reduction(&wg, windowSize, i, a)
        }
        wg.Wait()
        windowSize >>= 1
    }

    fmt.Printf("%d\n", a[0])
}
