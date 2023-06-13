package main

import (
    "fmt"
    "sync"
    "math/rand"
    "io"
    "os"
)

/*
Задание номер 7


Реализовать конкурентную запись данных в map.
*/

// Мапа
type SharedMap struct {
    mu sync.RWMutex
    shmap map[int]int
}


func (sh *SharedMap) Init() {
    sh.shmap = make(map[int]int)
}

// При записи мапа блокируется на запись с помощью мютекса
func (sh *SharedMap) Insert(key, value int) {
    sh.mu.Lock()
    defer sh.mu.Unlock()
    sh.shmap[key] = value
}

// При удалении мапа блокируется на запись с помощью мютекса
func (sh *SharedMap) Erase(key int) {
    sh.mu.Lock()
    defer sh.mu.Unlock()
    delete(sh.shmap, key)
}

// При удалении мапа блокируется на чтение с помощью мютекса
func (sh *SharedMap) Get(key int) (int, bool) {
    sh.mu.RLock()
    defer sh.mu.RUnlock()
    val, ok := sh.shmap[key]
    return val, ok
}


func (sh *SharedMap) Return(out io.Writer) {
    sh.mu.RLock()
    defer sh.mu.RUnlock()
    
    for k, v := range(sh.shmap) {
        fmt.Fprintf(out, "%6d : %6d\n", k, v)
    }
}


func worker(wg *sync.WaitGroup, shmap *SharedMap, id int, n int) {
    defer wg.Done()
    for i := 0; i < n; {
        
        k := rand.Intn(n * 100) + id * n * 100
        v := rand.Intn(10000)

        

        if _, ok := shmap.Get(k); !ok {
            shmap.Insert(k, v)
            i++
        }
    }
}


func main() {

    var n = 10
    var m = 9

    var shmap SharedMap
    shmap.Init()

    var wg sync.WaitGroup
    wg.Add(n)

    for i := 0; i < n; i++ {
        go worker(&wg, &shmap, i, m)
    }

    wg.Wait()
    
    shmap.Return(os.Stdin)
}
