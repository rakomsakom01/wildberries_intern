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
func reader(wg *sync.WaitGroup, id int, channel <-chan int) {
    defer wg.Done()
    for {
        num, ok := <-channel
        // проверка на закрытость канала
        if !ok {
            fmt.Printf("exit reader %d\n", id)
            return
        }

        fmt.Printf("%4d: %d\n", id, num)
    }
}

// принимает сигнал
func waitExit(wg *sync.WaitGroup, sig <-chan os.Signal, isExit *bool) {
    defer wg.Done()
    <-sig

    *isExit = true
}


func main() {
    // n - количество воркеров
    var n int

    fmt.Scanf("%d\n", &n)
    // канал для данных
    var channel = make(chan int)

    var wg sync.WaitGroup
    wg.Add(n + 1)

    for i := 0; i < n; i++ {
        go reader(&wg, i, channel)
    }
    
    // isExit - флаг того, что дана команда завершится
    var isExit = false
    var sig = make(chan os.Signal)
    // SIGINT - Ctrl + C
    signal.Notify(sig, syscall.SIGINT)

    go waitExit(&wg, sig, &isExit)

    // Отправляем данные в канал данных
    for {
        time.Sleep(time.Second)
        // Проверка, пришёл ли сигнал
        if isExit {
            // Извещаем воркеры закрытием канала
            close(channel)
            break
        }
        channel <- rand.Int()
    }
    
    // Ждём, пока все горутины завершатся
    wg.Wait()
}
