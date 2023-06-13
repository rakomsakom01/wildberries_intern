package main

import (
    "fmt"
    "sort"
)

/*
Задание номер 11

Реализовать пересечение двух неупорядоченных множеств.
*/


func main() {
    // n, m -количество элементов в множествах
    var n, m int

    fmt.Scanf("%d", &n)
    // a, b -множества
    var a = make([]int, n)

    for i := range(a) {
        fmt.Scanf("%d", &a[i])
    }

    fmt.Scanf("%d", &m)

    var b = make([]int, m)
    
    for i := range(b) {
        fmt.Scanf("%d", &b[i])
    }
    // Для начала отсортируем множества

    sort.Ints(a)
    sort.Ints(b)

    // Используем метод двух указателей
    // i привязан к a
    // j привязан к b
    var i = 0
    var j = 0

    var res = make([]int, 0, len(a) + len(b))

    for i < len(a) || j < len(b) {
        // Если одно из множеств закончилось, то мы добавляем в результат элементы другого множества
        if i == len(a) {
            res = append(res, b[j])
            j++
        } else if (j == len(b)) {
            res = append(res, a[i])
            i++
        } else if (a[i] < b[j]) {
            // В итоговое множество добавляем наименьший элемент
            res = append(res, a[i]) 
            i++
        } else if (a[i] > b[j]) {
            res = append(res, b[j])
            j++
        } else {
            // если наименьшие элементы совпадают, то объединяем
            res = append(res, a[i])
            i++
            j++
        }
    }

    fmt.Printf("%v\n", res)

}
