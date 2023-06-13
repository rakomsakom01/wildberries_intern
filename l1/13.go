package main

import (
    "fmt"
)

/*
Задание номер 13

Поменять местами два числа без создания временной переменной.
*/

func main() {
    var n, m int

    fmt.Scanf("%d %d", &n, &m)

    n = n + m // n = n + m, m = m
    m = n - m // n = n + m, m = n
    n = n - m // n = m, m = n


    fmt.Printf("%d %d\n", n, m)
}
