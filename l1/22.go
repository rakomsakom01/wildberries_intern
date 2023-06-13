package main

import (
    "fmt"
    "math/big"
)

/*
Задание номер 22

Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
    var b1, b2 big.Int

    fmt.Scanf("%s", &b1)
    fmt.Scanf("%s", &b2)

    var badd, bsub, bmul, bdiv, bmod big.Int

    // сложение
    badd.Add(&b1, &b2)
    // вычитание
    bsub.Sub(&b1, &b2)
    // умножение
    bmul.Mul(&b1, &b2)
    // деление
    bdiv.Div(&b1, &b2)
    // остаток
    bmod.Mod(&b1, &b2)
    

    fmt.Printf("%d \n%d \n%d \n%d \n%d\n", &badd, &bsub, &bmul, &bdiv, &bmod)
}
