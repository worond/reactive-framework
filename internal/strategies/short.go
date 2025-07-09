package strategies

import (
	"context"
	"reactive-framework/internal/dto"
	"reactive-framework/internal/services"
	"reactive-framework/internal/strategies/parts"
)

func Short(ctx context.Context) (*dto.Response, error) {
	user, location := parts.UserAndLocation(ctx)
	products := services.AsyncGetProducts(ctx, user, location)

	cart, err := services.AwaitGetCart(ctx, user, products, nil, nil)
	if err != nil {
		return nil, err
	}

	return &dto.Response{Cart: cart}, nil
}
