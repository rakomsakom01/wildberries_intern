package main

import (
    "fmt"
)


/*
Задание номер 12

Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
    // n -количество элементов
    var n int

    fmt.Scanf("%d", &n)
    
    // В качестве основы используем мапу
    var set = make(map[string]struct{})

    var a string
    for i := 0; i < n; i++ {
        fmt.Scanf("%s", &a)

        // проверка на наличии элемента в множестве
        _, ok := set[a]

        if !ok {
            set[a] = struct{}{}
        }
    }

    // вывод элементов в множестве
    for k, _ := range(set) {
        fmt.Printf("%s\n", k)
    }

}
