package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//用select实现一个生成随机数序列的程序
func main() {
    ch := make(chan int)
	//当有多个管道均可操作时，
	//select会随机选择一个管道。
    go func() {
        for {
            select {
            case ch <- 0:
            case ch <- 1:
            }
        }
    }()

    for v := range ch {
        fmt.Println(v)
    }
}

func worker(wg *sync.WaitGroup, cannel chan bool) {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println("helloworld")
        case <-cannel:
            return
        }
    }
}

func main1() {
    cancel := make(chan bool)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(&wg, cancel)
    }

    time.Sleep(time.Second)
    close(cancel)
    wg.Wait()
}

func worker2(ctx context.Context, wg *sync.WaitGroup) error {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println("helloworld")
        case <-ctx.Done():
            return ctx.Err()
        }
    }
}

func main2() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker2(ctx, &wg)
    }

    time.Sleep(time.Second)
    cancel()
    wg.Wait()
}