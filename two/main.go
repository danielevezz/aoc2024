package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)



func main() {
    count := ProcessInput()
    fmt.Println(count)
}

func IsSafe(l []int, withReductions bool) bool {
    var deltas []int
    var signs []int

    for i:=1; i<len(l);i++{
        delta := l[i] - l[i-1]
        if delta < -3 || delta > 3 || delta == 0 {
            return false
        }
        deltas = append(deltas, delta) 
        signs = append(signs, Sign(delta)) 
    }

    //fmt.Printf("%v\n",l)
    //fmt.Printf("%v\n",deltas)
    //fmt.Printf("%v\n\n",signs)
    
    isAllSame := AllSame(signs)
    

    if isAllSame {
        return true
    }
    
    if withReductions {
        return TryReductions(l)
    }

    return isAllSame

}

func TryReductions(l []int) bool {
    for _, reduction := range GenerateReductions(l) {
        if IsSafe(reduction,false) {
            return true 
        }
    }
    return false
}

func GenerateReductions(l []int) (reductions []([]int)) {
    fmt.Println(l)
    for i := range l {
        a := RemoveIndex(l,i)
        reductions = append(reductions,a)
    }
    return
}

func RemoveIndex(s []int, index int) []int {
    copy(s[index:], s[index+1:])
    return s[:len(s)-1]
}

func AllSame(a []int) bool {
    for i := 1; i < len(a); i++ {
        if a[i] != a[0] {
            return false
        }
    }
    return true
}

func Sign(i int) int {
    if i < 0 {
        return -1
    } else if i > 0 {
        return 1
    } else {
        return 0 
    }
}

func SplitLine(line string) []int {
    strs := strings.Split(line, " ")
    var ints []int
    for _, v := range strs {
        n, _ :=  strconv.Atoi(v)
        ints = append(ints,n) 
    }
    return ints
}

func ProcessInput() int {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    count := 0
    for scanner.Scan() {
        line := scanner.Text()
        if IsSafe(SplitLine(line),true) {
            count++
        }
    }
    return count
}
