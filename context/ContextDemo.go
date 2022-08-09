package main

import (
	"context"
	"fmt"
	"time"
)

/**
使用方法和设计原理 — 多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作
@see https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // 取消信号 ? 取消具体做了什么？

	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done(): // 订阅 ctx.Done() 管道中的消息 1
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done(): // 订阅 ctx.Done() 管道中的消息 2
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
