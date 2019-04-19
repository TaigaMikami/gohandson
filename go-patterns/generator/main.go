package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Count(ctx context.Context, start int, end int) chan int {
	ch := make(chan int)

	go func (ch chan<- int) {
		defer close(ch)
	loop:
		for i := start; i<= end; i++ {
			select {
			case <-ctx.Done():
				break loop
			default:
			}
			// 重い処理
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
	}(ch)
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// ctr + c でキャンセルする
	canceled := false
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		cancel()
		canceled = true
	}()

	for i := range Count(ctx,1, 99) {
		fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
	}
	fmt.Println(100, "bottles of beer on the wall")
	if canceled {
		fmt.Fprintln(os.Stderr, "canceled")
	}
}
