package main

import (
    "fmt"
    "math/rand"
    "os"
    "os/signal"
    "syscall"
    "time"
    "sync"
)

/*
Задание номер 4

Реализовать постоянную запись данных в канал (главный поток). 
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

*/

// Просто читает данные из канала
func reader(wg *sync.WaitGroup, id int, channel <-chan int, exit <-chan bool) {
    defer wg.Done()
    for {
        select {
        // считываем данные из канала
        case num := <-channel:
            fmt.Printf("%4d: %d\n", id, num)
        // если пришёл сигнал о завершении
        case <- exit:
            fmt.Printf("exit reader %d\n", id)
            return
        }
    }
}

// принимает сигнал
func waitExit(wg *sync.WaitGroup, sig <-chan os.Signal, n int, exit chan<- bool) {
    defer wg.Done()
    <-sig
    // отправляем сигналы о завершении каждому воркеру и главному потоку
    for i := 0; i < n + 1; i++ {
        exit <- true
    }
}


func main() {
    var n int

    fmt.Scanf("%d\n", &n)
    // канал для данных
    var channel = make(chan int)
    // канал для сигналов о завершении
    var exit = make(chan bool)

    var wg sync.WaitGroup
    wg.Add(n + 1)

    for i := 0; i < n; i++ {
        go reader(&wg, i, channel, exit)
    }
    
    var sig = make(chan os.Signal)
    // SIGINT - Ctrl + C
    signal.Notify(sig, syscall.SIGINT)

    go waitExit(&wg, sig, n, exit)

    // Отправляем данные в канал данных
    mainfor:
    for {
        time.Sleep(time.Second)
        
        select {
        // если пришёл сигнал о завершении
        case <- exit:
            break mainfor
        // иначе пишем данные в канл данных
        default:
            channel <- rand.Int()
        }
    }

    // Ждём, пока все горутины завершатся
    wg.Wait()
}
