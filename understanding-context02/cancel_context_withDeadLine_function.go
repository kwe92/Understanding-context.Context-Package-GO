package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: review code and comments

// TODO: add comments for context.WithDeadLine

func main() {
	// instatiate the main parent context to be passed throughout your program
	ctx := context.Background()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	// instatiate a new time.Time object and add the amount of time before cancelation
	var deadLine time.Time = time.Now().Add(3 * time.Second)

	// create a new context.Context scoped locally to the doSomething
	// and a cancelation function returned by context.WithDeadline
	// context.WithDeadline takes two arguments: parent context and a time.Time object representing the time till cancelation
	ctx, cancelCtx := context.WithDeadline(ctx, deadLine)

	// defer the cancelation of the context if the deadline is not exceeded
	defer cancelCtx()

	// create out channel of integers to pass to worker goroutine
	var out chan int = make(chan int)

	// launch a worker go routine that does something with the out channel passed in
	go doAnother(ctx, out)

	// start writting n number of values to the out channel
	for n := 0; n < 5; n++ {
		// use select statement to check if channel is closed
		// why is that needed here?
		// because go routines will deadlock otherwise as the doAnother will continue trying to send number to a closed channel

		select {
		// write the nth integer to out channel
		case out <- n:

			// wait 1 second between channel writes
			time.Sleep(1 * time.Second)

		case <-ctx.Done():
			// break is used to break out of the for loop and continue | if return was used the function would immediately end
			break
		}
	}

	time.Sleep(200 * time.Millisecond)

	fmt.Println("doSomething has finished.")

}

// doAnother: acts as a worker go routine
func doAnother(ctx context.Context, in <-chan int) {
	for {
		select {

		// receive the nth value ready to be read from the in channel
		case receivedChannelValue := <-in:

			// process the nth value
			fmt.Println("value received from in channel:", receivedChannelValue)

		// check if context is canceled
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Println("context error in doAnother go routine:", ctx.Err())
			}
			fmt.Println("doAnother has finished due to:", ctx.Err())

			// end the goroutine
			return

		// periodically indicate that there are no channel read / write operations ready and the context is not canceled
		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("There are no available channel read/write operations and the context is still open...")

		}
	}
}

// context.WithDeadLine
