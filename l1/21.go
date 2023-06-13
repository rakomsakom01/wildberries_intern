package main

import (
    "fmt"
)

/*
Задание номер 21

Реализовать паттерн «адаптер» на любом примере.
*/

// Исходный класс
type WWLegacy struct {
    name1 string
    name2 string
    name3 string
    name4 string
}


func (ww *WWLegacy) Init() {
    ww.name1 = "Arctic Burn"
    ww.name2 = "Splinter Blast"
    ww.name3 = "Cold Embrace"
    ww.name4 = "Winter's Curse"
}


func (ww *WWLegacy) R () { 
    fmt.Printf("Cast %s!\n", ww.name1)
}


func (ww *WWLegacy) T () { 
    fmt.Printf("Cast %s!\n", ww.name2)
}


func (ww *WWLegacy) E () { 
    fmt.Printf("Cast %s!\n", ww.name3)
}


func (ww *WWLegacy) W () { 
    fmt.Printf("Cast %s!\n", ww.name4)
}


// интерфейс
type Hero interface {
    Q()
    W()
    E()
    R()
}

// основная функция который работает с интерфейсом
func CastSpells(hero Hero) {
    hero.Q()
    hero.W()
    hero.E()
    hero.R()
}

// Адаптер
type WWAdapter struct {
    ww *WWLegacy
}


func (ww *WWAdapter) Init(wwl *WWLegacy) {
    ww.ww = wwl
}


func (ww *WWAdapter) Q() {
    ww.ww.R()
}


func (ww *WWAdapter) W() {
    ww.ww.T()
}


func (ww *WWAdapter) E() {
    ww.ww.E()
}


func (ww *WWAdapter) R() {
    ww.ww.W()
}


func main() {
    var wwl WWLegacy
    wwl.Init()
    var ww WWAdapter
    ww.Init(&wwl)

    CastSpells(&ww)
    
}
