package strategies

import (
	"context"
	"reactive-framework/internal/dto"
	"reactive-framework/internal/promise"
	"reactive-framework/internal/services"
	"reactive-framework/internal/strategies/parts"
)

func Short(ctx context.Context) (*dto.Response, error) {
	user, location := parts.UserAndLocation(ctx)
	products := services.AsyncGetProducts(ctx, user, location)
	cart := services.AsyncGetCart(ctx, user, products, nil, nil)

	result, err := promise.Await(cart)
	if err != nil {
		return nil, err
	}

	return &dto.Response{CartResponse: result}, nil
}
