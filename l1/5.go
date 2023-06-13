package main

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
    "flag"
)

/*
Задание номер 5

Разработать программу, которая будет последовательно отправлять значения в канал, 
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

// Просто читает данные из канала
func reader(wg *sync.WaitGroup, channel <-chan int, exit <-chan bool) {
    defer wg.Done()
    for {
        select {
        // считываем данные из канала
        case num := <-channel:
            fmt.Printf("%d\n", num)
        // если пришёл сигнал о завершении
        case <- exit:
            fmt.Printf("exit reader\n")
            return
        }
    }
}


// Просто записывает данные в канал
func writer(wg *sync.WaitGroup, channel chan<- int, exit <-chan bool) {
    defer wg.Done()

    var tick = time.Tick(400 * time.Millisecond)

    for {
        select {
        // записываем данные в канал
        case <-tick:
            channel <- rand.Int()
        // если пришёл сигнал о завершении
        case <-exit: 
            fmt.Printf("exit writer\n")
            return
        }
    }   
}


// Завершает работу по истечению n секунд
func waitExit(wg *sync.WaitGroup, exit chan<- bool, waitTime <-chan time.Time) {
    defer wg.Done()

    <-waitTime
    // отправляем сигналы о завершении каждому воркеру
    for i := 0; i < 2; i++ {
        exit <- true
    }
}


func main() {
    // n - количество ворекров

    // n передаётся в аргументах для более точного замера времени работы
    var n = flag.Int("t", 2, "Timer")
    flag.Parse()

    // Запкскаем таймер
    var waitTime = time.After(time.Duration(*n) * time.Second)

    // канал для данных
    var channel = make(chan int)
    // канал для сигналов о завершении
    var exit = make(chan bool)

    var wg sync.WaitGroup
    wg.Add(3)

    go reader(&wg, channel, exit)
    go writer(&wg, channel, exit)
    // запускаем таймер
    go waitExit(&wg, exit, waitTime)

    // Ждём, пока все горутины завершатся
    wg.Wait()
}
