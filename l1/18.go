package main

import (
    "fmt"
    "sync"
)

/*
Задание номер 18

Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. 
По завершению программа должна выводить итоговое значение счетчика.
*/


// Счётчик
type Clicker struct{
    sync.RWMutex
    click int
}


func (c *Clicker) Click() {
    c.Lock()
    c.click++
    c.Unlock()
}


func (c *Clicker) Value() int {
    var ret int
    c.RLock()
    ret = c.click
    c.RUnlock()
    return ret
}


func workerClicker(wg *sync.WaitGroup, c *Clicker, n int) {
    defer wg.Done()    

    for i := 0; i < n; i++ {
        c.Click()
    }
}


func main() {
    // w - колчество вркеров
    // n - количество кликов от одного воркера
    var w, n int

    fmt.Scanf("%d %d", &w, &n)


    var wg sync.WaitGroup
    var clicker Clicker

    wg.Add(w)

    for i := 0; i < w; i++ {
        go workerClicker(&wg, &clicker, n)
    }
    
    wg.Wait()

    fmt.Printf("%d\n", clicker.Value())
}
