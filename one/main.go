package main

import (
    "container/heap"
    "math"
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
    "strings"
)

func main() {

    l1 := &IntHeap{}
    heap.Init(l1)
    l2 := &IntHeap{}
    heap.Init(l2)

    GetLists(l1, l2)

    part2(l1,l2)
    //part1(l1,l2)

}

func part1(l1, l2 *IntHeap) {
    sum := 0.0
    for l1.Len() > 0 {

        n1 := heap.Pop(l1).(int)
        n2 := heap.Pop(l2).(int)

        fmt.Println(n1, n2)

        sum = sum + math.Abs(float64(n1)- float64(n2))
    }

    fmt.Printf("%f",sum)
}

func part2(l1, l2 *IntHeap) {

    sum := 0
    for l1.Len() > 0 {
        f1, f2 := readUntilChange(l1)

        f3 := countNumbers(l2, f1)

        fmt.Printf("value1: %d, count1: %d, count2: %d\n", f1,f2,f3)

        sum += f1 * f2 * f3
    }

    fmt.Println(sum)
}

func readUntilChange(h *IntHeap) (value int, count int) {
    value = (*h)[0]
    count = 1

    for h.Len() > 0 {
        n := heap.Pop(h).(int)
        if h.Len() <= 0 || n != (*h)[0] {
            return
        }
        count++
    }
    return
}

func countNumbers(h *IntHeap, v int) (count int) {
    for i:=0; i<h.Len();i++ {
        n := (*h)[i]
        if n == v {
            count++
        }
    }
    return
}


// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func GetLists(l1, l2 *IntHeap) {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        splitInsert(scanner.Text(),l1,l2)
    }
}

func splitInsert(line string, l1, l2 *IntHeap) {
    numbers := strings.Split(line, "   ")

    n, _ := strconv.Atoi(numbers[0])
    heap.Push(l1, n)

    n, _ = strconv.Atoi(numbers[1])
    heap.Push(l2, n)
}
