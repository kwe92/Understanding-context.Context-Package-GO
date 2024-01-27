package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// instatiate parent context
	ctx := context.Background()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {

	// specify context timeout
	var timeOut time.Duration = 3 * time.Second

	ctx, cancelContext := context.WithTimeout(ctx, timeOut)

	defer cancelContext()

	var out chan int = make(chan int)

	go doAnother(ctx, out)

	for n := 0; n < 5; n++ {
		select {
		case out <- n:
			time.Sleep(1 * time.Second)

		case <-ctx.Done():
			break
		}
	}

	time.Sleep(200 * time.Millisecond)

	fmt.Println("doSomething has finished doing the stuff.")
}

func doAnother(ctx context.Context, in <-chan int) {
	for {
		select {
		case receivedChannelValue := <-in:
			fmt.Println("value recieved from in channel:", receivedChannelValue)

		case <-ctx.Done():
			var err error

			if err = ctx.Err(); err != nil {
				fmt.Println("error from context:", err.Error())
			}

			fmt.Println("doAnother finished due to:", err)
			// end the goroutine
			return

		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("waiting for available channel read/write operation or the context to be canceled...")
		}
	}
}

// context.WithTimeOut

//   - similar to context.DeadLine with the exception that instead of using a
//     time.Time for context expiration you use a time.Duration which adds brevity
