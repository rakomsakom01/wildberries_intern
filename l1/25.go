package main

import (
    _"fmt"
    "time"
)
/*
Задание номер 25

Реализовать собственную функцию sleep.
*/

func Sleep(d time.Duration) {
    // trigger - канал, After посылает данные в канал через время d    
    var trigger = time.After(d)
    // Ждём
    <-trigger
}



func main() {
    //time.Sleep(10 * time.Second)
    Sleep(10 * time.Second)
}
