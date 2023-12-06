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

	// pass context to function
	doSomething(ctx)
}

// doSomething: sends work to worker go routines
func doSomething(ctx context.Context) {

	// assign new context with associated cancelation function within scope of doSomething
	ctx, cancelCtx := context.WithCancel(ctx)

	// create out channel of integers to pass to worker goroutine
	out := make(chan int)

	// run worker goroutine doing something with out channel values
	go doAnother(ctx, out)

	// write n number of integers to out channel
	// for-loop required to write to out channel
	// if there are no values written to the out channel then there will be nothing to read
	for n := 0; n < 3; n++ {

		// wait for 200 miliseconds before writing to out channel
		time.Sleep(200 * time.Millisecond)

		// write n to out channel
		out <- n
	}

	// cancel context after for loop runs
	cancelCtx()

	// sleep for 1 second before finishing
	time.Sleep(1 * time.Second)

	fmt.Println("doSomething finished.")
}

// doAnother: acts as a worker go routine
func doAnother(ctx context.Context, in <-chan int) {

	for {
		select {

		// read from in channel if there are values to be read
		case num := <-in:
			// process value doing something with it
			fmt.Println("value recieved from in channel:", num)

		// cancelation value does not need to be assigned to variable
		// nil return values are skipped by for select statements
		// an empty struct is the value returned by a closed channel via cancelation as they take up no memory
		case cancelChannelValue := <-ctx.Done():
			// when canceling a context.Context an error `context canceled` will be generated
			// retrieved by calling context.Err()
			if err := ctx.Err(); err != nil {
				fmt.Println("error in doAnother:", err.Error())
			}

			fmt.Println("doAnother finished. canceled channel value:", cancelChannelValue)

			// end the goroutine
			return

		// if there are no values to be read and the channel is not canceled the default statement runs
		default:
			// Periodically run the defaut channel as not to take up too much memory if there are no channel read / write operations ready
			time.Sleep(100 * time.Millisecond)
			fmt.Println("waiting for channel read/write operation or context cancelation...")

		}
	}

}

// Channel Variable Names

//   - out: write-channel not explicitly stated
//   - in: read-only channel / receiver channel explicitly stated within parameter definition
