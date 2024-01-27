package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// TODO: Review why error custom error is not printing out properly
// TODO: implement context.WithTimeoutCause

func main() {
	ctx := context.Background()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {

	deadline := time.Now().Add(3 * time.Second)

	customDeadlineCancellationError := errors.New("doAnother RPC timeout")

	ctx, cancelCtx := context.WithDeadlineCause(ctx, deadline, customDeadlineCancellationError)

	defer cancelCtx()

	out := make(chan int)

	doAnother(ctx, out)

	func() {
		for n := 0; n < 5; n++ {
			select {
			case out <- n:

			case <-ctx.Done():
				fmt.Println("context has been closed due to:", ctx.Err())
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("doSomething has finished.")
}

func doAnother(ctx context.Context, in <-chan int) {

	go func() {
		for {
			select {
			case resultFromInChannel := <-in:
				fmt.Println("processed result from in channel:", resultFromInChannel+42)
				time.Sleep(4 * time.Second)

			case <-ctx.Done():
				return
			}
		}
	}()
}

// context.WithDeadlineCause
