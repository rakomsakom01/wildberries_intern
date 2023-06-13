package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

/*
Задание номер 20

Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow».
*/


func reverse(str string) string {
    // Разбиваем строку на слова
    words := strings.Split(str, " ")

    reverseWords := make([]string, len(words))

    for i := range(words) {
        reverseWords[len(words) - 1 - i] = words[i]
    }

    // склеиваем строку
    return strings.Join(reverseWords, " ")
}


func main() {
    var s string

    reader := bufio.NewReader(os.Stdin)

    s, _ = reader.ReadString('\n')

    fmt.Printf("%s\n", reverse(s))

}
