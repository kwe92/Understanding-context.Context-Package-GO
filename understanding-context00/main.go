package main

import (
	"context"
	"fmt"
)

// TODO: review code and edit comments

// globally accessible context.Context keys
const globalContextKey0 = "someStoredValue0"
const globalContextKey1 = "someStoredValue1"

func main() {
	// ctx := context.TODO()

	// instatiate parent context
	ctx := context.Background()

	// add key/value pair to context | returns new context.Context instead of mutating original
	ctx = context.WithValue(ctx, globalContextKey0, 42)

	doSomething(ctx)

	// attempt to return a value that was added by the doSomething function from the context.Context passed into it
	uncapturedValue := ctx.Value(globalContextKey1)

	fmt.Printf("uncaptured value: %v\n", uncapturedValue)

}

func doSomething(ctx context.Context) {

	fmt.Printf("doSomething context value: %v\n", ctx.Value(globalContextKey0))

	// context.Context passed by value not by reference/pointer
	ctx = context.WithValue(ctx, globalContextKey1, "I WILL NOT BE SAVED IN PARENT CONTEXT!")
}

// What is context.Context?

//   - a package in GO part of the standard library
//   - supplies functions with additional information
//   - assists with cancellation and graceful termination,
//     saving computing resources when operations are long running or
//     their value will not be processed for some reason
//   - can store and retrieve data in O(1) constant-time
//   - can signal to other operations that a process is finsished
//     which can be extremely helpful when tasks are ran concurrently

// Why is context.Context Useful?

//   - timeout long running operations
//   - gracefully shutdown servers and databases
//     by propagating cancellation signals and cleaning up tasks
//   - propagation of cancelation signals for concurrently running tasks
//   - cancellation can save computing resources on a busy server
//   - control cancellation, propagation and life cycle of requests
//   - stop long running operations and database calls
//   - stop processes that will not be received due to client disconnection

// context.TODO()

//   - instantiate a generic context used
//     when you dont know which context to use
//   - should be replaced once you know which context to use

// context.Background()

//   - does the exact same thing as context.TODO()
//     but indicates to developers that the context will be used

// Functions With context.Context as a Parameter

//   - it is idomatic that functions with context.Context as a parameter define
//     the context.Context variable as the first parameter in the function head

// Store and Accessing Values Within a Context

//   - values are stored as key/value pairs O(1) operations
//     using the context.WithValue function

// context.WithValue

//   - adds key/value pairs to a given context
//   - has three parameters:
//       - the parent context.Context i.e. the context.Context you wish to add the key/value pair to
//       - a key of any type
//       - a value of any type
//   - returns a new context.Context with all previous data preserved

// context.Context keys

//   - typically stored in constant global variables
//     so the propagated context can have its values accessed
