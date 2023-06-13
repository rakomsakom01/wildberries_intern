package main

import (
    "fmt"
    "sort"
    "math/rand"
)

/*
Задание 16

Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/


func quickSort(a []int) {
    quickSortPart(a, 0, len(a) - 1) 
}

// Сортит часть от left до right
func quickSortPart(a []int, left int, right int) {
    if left < right {
        p := partition(a, left, right)

        quickSortPart(a, left, p)
        quickSortPart(a, p + 1, right)
    }
}

// возвращает индекс опорного элемента и перераспределяет элементы в массиве таким образом, 
// что элементы, меньшие опорного, помещаются перед ним, а большие или равные - после.
func partition(a []int, left int, right int) int {
    var pval int
    pval = a[(left + right) / 2]

    var i = left
    var j = right

    for {
        for i < right && a[i] < pval {
            i++
        }
        for j > left && a[j] > pval {
            j--
        }
        if i >= j {
            break
        }
                
        a[i], a[j] = a[j], a[i]
        i++
        j--
    }

    return j
}


func compareArrays(a, b []int) bool {
    if (len(a) != len(b)) {
        return false
    }

    for i:= range(a) {
        if (a[i] != b[i]) {
            return false
        }
    }

    return true
}

// стресс тест и сравнение с стандартной реализацией сортировки
func stressTest(iters int, lenMin int, lenMax int, valMin int, valMax int) {
    for i:= 0; i < iters; i++ {
        var lenArr = rand.Intn(lenMax - lenMin) + lenMin

        var rawArr = make([]int, lenArr)

        for i := range(rawArr) {
            rawArr[i] = rand.Intn(valMax - valMin) + valMin
        }

        var test1 = make([]int, lenArr)
        var test2 = make([]int, lenArr)
        copy(test1, rawArr)
        copy(test2, rawArr)
        quickSort(test1)
        sort.Ints(test2)

        if !compareArrays(test1, test2) {
            fmt.Println(rawArr)
            fmt.Println(test1)
            fmt.Println(test2)
        }
    }
}



func main() {
    stressTest(100, 1000, 20000, 1, 10)
}
