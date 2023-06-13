package main

import (
    _"fmt"
    "math/rand"
)



/*
Задание номер 15

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? 

Приведите корректный пример реализации.

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}

func main() {
    someFunc()
}
*/

/*

Ответ: Если строка от createHugeString содержит символы не из латинского алфавита(например, 
из арабского, китайского или русского алфавита), то при попытке получить подстроку длины n от такой 
строки мы получим строку длиной меньше, чем n, причём есть вероятность того, что в конце такой строки
будет мусор, который будет некорректно отображатся.

Это связано с тем, что строка это массив byte(1 байт), а некоторые символы могут кодироваться 2 
и более байтами.

Чтобы не допустить такого, необходимо преобразовать исходную строку в массив rune(4 бита), причём при
преобразовании будет учтен разный размер символов.
*/

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    // вот тута
    justString = string([]rune(v)[:100])
}

func main() {
    someFunc()
}