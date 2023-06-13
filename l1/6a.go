package main

import (
    "fmt"
    "sync"
)

/*
Задание номер 6

Реализовать все возможные способы остановки выполнения горутины. 
*/

// воторй способ - через канал, закрыв его
func worker(wg *sync.WaitGroup, channel <-chan int) {
    defer wg.Done()
    for {
        val, ok := <-channel 

        if !ok {
            return
        }

        fmt.Printf("%d\n", val)
    }
}


func main() {
    var wg sync.WaitGroup 

    wg.Add(1)

    var channel = make(chan int)

    go worker(&wg, channel)

    channel <- 1
    channel <- 2

    close(channel)

    wg.Wait()
}
