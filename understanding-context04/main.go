package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: Refactor to return channels from go routines | maybe add wait groups
// TODO: add comments

func main() {

	ctx := context.Background()

	doSomething(ctx)

}

func doSomething(ctx context.Context) {

	timeOut := 3 * time.Second

	ctx, cancelCtx := context.WithTimeout(ctx, timeOut)

	defer cancelCtx()

	out := make(chan int)

	doAnother(ctx, out)

	doSomethingAfter(ctx, out)

	for n := 0; n < 12; n++ {

		select {
		case out <- n:
			fmt.Println("value written to out channel:", n)
			time.Sleep(1 * time.Second)

		case <-ctx.Done():
			break
		}
	}

	time.Sleep(200 * time.Millisecond)

	fmt.Println("doSomething has finished processing.")

}

func doAnother(ctx context.Context, in <-chan int) {

	go func() {
		for {
			select {
			case receivedChannelValue := <-in:
				fmt.Println("processed received channel value:", receivedChannelValue*2)

			case <-ctx.Done():
				fmt.Println("from doAnother function: context has been closed due to:", ctx.Err().Error())
				return
			default:
				time.Sleep(200 * time.Millisecond)
				fmt.Println("waiting for channel read/write operation or Context cancellation...")
			}
		}
	}()

}

func doSomethingAfter(ctx context.Context, in <-chan int) {
	context.AfterFunc(ctx, func() {
		fmt.Println("additional processing after context is canceled")
	})

	go func() {
		for {
			select {
			case receivedChannelValue := <-in:
				fmt.Println("processed received channel value in doSomethingAfter function:", receivedChannelValue*3)
			case <-ctx.Done():
				fmt.Println("from doSomethingAfter function: context has been closed due to:", ctx.Err().Error())
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("doSomethingAfter waiting for channel read/write operation or context cancellation......")
			}
		}
	}()
}

// context.AfterFunc
