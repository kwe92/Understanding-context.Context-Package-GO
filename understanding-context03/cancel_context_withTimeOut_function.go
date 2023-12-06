package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: review code and comments

func main() {
	// instatiate parent context
	ctx := context.Background()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {

	// specify context timeout
	var timeOut time.Duration = 3 * time.Second

	// create a new context.Context scoped locally to the doSomething
	// and a cancelation function returned by context.WithDeadline
	// context.WithDeadline takes two arguments: parent context and a time.Duration object representing the time till cancelation
	ctx, cancelContext := context.WithTimeout(ctx, timeOut)

	// cancel context if timeout has not been exceeded
	defer cancelContext()

	// create out channel of integers to pass to worker goroutine
	var out chan int = make(chan int)

	go doAnother(ctx, out)

	for n := 0; n < 5; n++ {
		select {
		// write the nth integer to out channel
		case out <- n:
			// wait for 1 second between out channel writes
			time.Sleep(1 * time.Second)

		// if the context is done stop attempting to write values to out channel and break out of the loop continuing function execution
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
		// read from in channel
		case receivedChannelValue := <-in:
			// process in channel value
			fmt.Println("value recieved from in channel:", receivedChannelValue)

		case <-ctx.Done():
			var err error

			if err = ctx.Err(); err != nil {
				// TODO: send custom error depending on the error recieved | maybe see what the error string contains
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
