package main

import (
    "fmt"
    "sync"
)

/*
Задание номер 6

Реализовать все возможные способы остановки выполнения горутины. 
*/

// первый способ - через канал, используя отдельный для извещения о остановки
func worker(wg *sync.WaitGroup, channel <-chan int, stop <-chan bool) {
    defer wg.Done()
    for {
        select {
        case val := <-channel:
            fmt.Printf("%d\n", val)
        case <-stop:
            return
        }
    }
}


func main() {
    var wg sync.WaitGroup 

    wg.Add(1)

    var channel = make(chan int)
    var stop = make(chan bool)

    go worker(&wg, channel, stop)

    channel <- 1
    channel <- 2

    stop <- true
    wg.Wait()
}
