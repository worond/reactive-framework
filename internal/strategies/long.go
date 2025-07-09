package strategies

import (
	"context"
	"reactive-framework/internal/dto"
	"reactive-framework/internal/services"
)

func Long(ctx context.Context) (*dto.Response, error) {
	user := services.AsyncGetUser(ctx)
	location := services.AsyncGetLocation(ctx)
	products := services.AsyncGetProducts(ctx, user, location)
	prices := services.AsyncGetPrices(ctx, user, products)
	labels := services.AsyncGetLabels(ctx, products)

	cart, err := services.AwaitGetCart(ctx, user, products, prices, labels)
	if err != nil {
		return nil, err
	}

	return &dto.Response{Cart: cart}, nil
}
