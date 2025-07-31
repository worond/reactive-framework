package strategies

import (
	"context"
	"fmt"
	"reactive-framework/internal/dto"
	"reactive-framework/internal/promise"
	"reactive-framework/internal/services"
	"time"
)

func Long(ctx context.Context) (*dto.Response, error) {
	start := time.Now()

	location := services.AsyncGetLocation(ctx)                         // 0ms
	user := services.AsyncGetUser(ctx)                                 // 0ms
	products := services.AsyncGetProducts(ctx, user, location)         // 100ms, 200ms = 200ms
	labels := services.AsyncGetLabels(ctx, products)                   // 200ms + 100ms = 300ms
	prices := services.AsyncGetPrices(ctx, user, products)             // 100ms, 200ms + 100ms = 300ms
	cart := services.AsyncGetCart(ctx, user, products, prices, labels) // 100ms, 200ms + 100ms, 300ms + 100ms, 300ms + 100ms = 400ms

	fmt.Printf("Async funcs took: %s\n", time.Since(start))

	result, err := promise.Await(cart) // 400ms = 500ms
	if err != nil {
		return nil, err
	}

	fmt.Printf("Await took: %s\n", time.Since(start))

	return &dto.Response{CartResponse: result}, nil
}
