package main

import (
    "adventofcode24/one"
    "container/heap"
    "math"
    "fmt"
)

func main() {
    
    l1 := &one.IntHeap{}
    heap.Init(l1)
    l2 := &one.IntHeap{}
    heap.Init(l2)

    one.GetLists(l1, l2)

    sum := 0.0

    for l1.Len() > 0 {
        n1 := heap.Pop(l1).(int)
        n2 := heap.Pop(l2).(int)

        

        fmt.Println(n1, n2)
        sum = sum + math.Abs(float64(n1)- float64(n2))
    }

    fmt.Printf("%f",sum)
}
