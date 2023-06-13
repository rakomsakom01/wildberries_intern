package main

import (
    "sync"
    "fmt"
)

/*
Задание номер 9

Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

*/

// Воркер, который читает данные из массива и отправляет их в канал
func workerInput(wg *sync.WaitGroup, rawData chan<- int, a []int) {
    defer wg.Done()
    for i := range(a) {
        rawData <- a[i]    
    }ё
    close(rawData)
}

// Воркер, который читает данные из одного канала, перерабатывает и отправляет их в другой канал
func workerPipeline(wg *sync.WaitGroup, rawData <-chan int, readyData chan<- int) {
    defer wg.Done()
    for {
        raw, ok := <-rawData
        if !ok {
            close(readyData)
            return
        }
        readyData <- raw * 2
    }
}

// Воркер, который выводит данные из канала в stdout
func workerOutput(wg *sync.WaitGroup, readyData <-chan int) {
    defer wg.Done()
    for {
        data, ok := <-readyData

        if !ok {
            return
        }

        fmt.Printf("%d\n", data)
    }
}


func main() {
    // n - длина массива
    var n int

    fmt.Scanf("%d", &n)
    // a - массив 
    var a = make([]int, n)
    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }

    var wg sync.WaitGroup
    wg.Add(3)

    rawData := make(chan int)

    readyData := make(chan int)

    go workerInput(&wg, rawData, a)
    go workerPipeline(&wg, rawData, readyData)
    go workerOutput(&wg, readyData)

    wg.Wait()
    
}
