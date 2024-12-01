package one

import (
    "log"
    "bufio"
    "os"
    "strings"
    "strconv"
    "container/heap"
)

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
    file, err := os.Open("one/input.txt")
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
