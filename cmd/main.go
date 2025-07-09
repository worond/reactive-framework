package main

import (
	"context"
	"fmt"
	"reactive-framework/internal/strategies"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	start := time.Now()
	res, err := strategies.Long(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Long strategies took: %s\n", time.Since(start))
	fmt.Printf("Long response: %+v\n\n", res)

	start = time.Now()
	res, err = strategies.Short(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Short strategies took: %s\n", time.Since(start))
	fmt.Printf("Short response: %+v\n\n", res)

	start = time.Now()
	res, err = strategies.Degraded(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Degraded strategies took: %s\n", time.Since(start))
	fmt.Printf("Degraded response: %+v\n\n", res)

}
