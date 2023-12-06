package main

import (
	"context"
	"fmt"
)

// TODO: review code and comments

// globally accessible context.Context key
const globalContextKey0 = "someStoredValue0"
const globalContextKey1 = "someStoredValue1"

func main() {
	// ctx := context.TODO()

	// instatiate parent context
	ctx := context.Background()

	// add key/value pair to context | returns new context.Context instead of mutating original
	ctx = context.WithValue(ctx, globalContextKey0, 43)

	doSomething(ctx)

	// attempt to return a value that was added by the doSomething function from the context.Context passed into it
	uncapturedValue := ctx.Value(globalContextKey1)

	fmt.Printf("uncaptured value: %v\n", uncapturedValue)

}

func doSomething(ctx context.Context) {

	// fmt.Println("Do something!!!!")
	fmt.Printf("doSomething context value: %v\n", ctx.Value(globalContextKey0))

	// context.Context passed by value not by reference/pointer
	ctx = context.WithValue(ctx, globalContextKey1, "I WILL NOT BE SAVED IN PARENT SCOPE!")
}

// context.TODO()

//   - instantiate a generic context
//     when you dont know which context to use

// context.Background()

//   - does the exact same thing as context.TODO()
//     but indicates to developers that this context will be used

// Functions With Context as a Parameter

//   - it is idomatic that functions with context.Context as a parameter should ensure that
//     the context.Context is the first parameter in the function

//   - it is idomatic that context.Context is the first parameter
//     in functions with context.Context as a parameter

// Store and Accessing Values Within a Context

//   - values are stored as key/value pairs O(1) access

// context.WithValue

//   - adds key/value pairs to a given context
//   - has three parameters
//       - the parent context.Context i.e. the context.Context you wish to add the key/value pair to
//       - the key of any type
//       - the value of any type
//   - returns a new context.Context with all previous data preserved

// context.Context keys

//   - typically stored in global variables
