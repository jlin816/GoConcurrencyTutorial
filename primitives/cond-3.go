package main

import (
    "sync"
    "fmt"
    "time"
    "math/rand"
)

func main() {
    var buf []int
    var mu sync.Mutex
    cond := sync.NewCond(&mu)

    go producer(&buf, &mu, cond)

    for {
        mu.Lock()
        for len(buf) == 0 {
            cond.Wait()
        }
        for _, v := range buf {
            process(v)
        }
        buf = nil // empty slice
        mu.Unlock()
    }
}

func process(v int) {
    fmt.Printf("processing %d\n", v)
}

func producer(buf *[]int, mu *sync.Mutex, cond *sync.Cond) {
    for {
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        mu.Lock()
        *buf = append(*buf, rand.Intn(1000))
        cond.Broadcast()
        mu.Unlock()
    }
}
