package strategies

import (
	"context"
	"reactive-framework/internal/dto"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
	"reactive-framework/internal/services"
)

func Degraded(ctx context.Context) (*dto.Response, error) {
	user := services.AsyncGetUser(ctx)
	location := services.AsyncGetLocation(ctx)
	products := services.AsyncGetProducts(ctx, user, location)
	prices := services.AsyncGetPrices(ctx, user, products)
	labels := services.AsyncGetLabelsWithError(ctx, products).Degradable()
	cart := services.AsyncGetCart(ctx, user, products, prices, labels)

	result, err := promise.Await[entities.Cart](cart)
	if err != nil {
		return nil, err
	}

	return &dto.Response{CartResponse: result}, nil
}
