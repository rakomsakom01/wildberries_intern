package main

import (
    "fmt"
    "bufio"
    "os"
)

/*
Задание номер 26

Разработать программу, которая проверяет, что все символы в строке уникальные 
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
*/

// приводит в нижний регистр
func convert(c rune) rune {
    if (c >= 'A' && c <= 'Z') {
        return c - 'A' + 'a'
    }
    return c
}

func isUnique(str string) bool {
    // здесь храним все символы, кторые сожержатся в исходной строке
    var set = make(map[rune]struct{})

    for _, c := range(str) {
        // проверка на уникальность
        rightChar := convert(c)
        _, ok := set[rightChar]

        if ok {
            return false
        }
        set[rightChar] = struct{}{}
    }

    return true
}


func main() {
    var s string

    reader := bufio.NewReader(os.Stdin)

    s, _ = reader.ReadString('\n')

    fmt.Printf("%t\n", isUnique(s))

}
