package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
	"time"
)

func getPrices(user entities.User, products entities.Products) (entities.Prices, error) {
	time.Sleep(100 * time.Millisecond)

	prices := make(entities.Prices, len(products))
	for _, product := range products {
		prices[product.ID] = entities.Price{
			Original:   float64(product.Quantity * 1000),
			Discounted: float64(product.Quantity * 100),
		}
	}

	return prices, nil
}

func AsyncGetPrices(
	ctx context.Context,
	userPromise *observer.Promise[entities.User],
	productsPromise *observer.Promise[entities.Products],
) *observer.Promise[entities.Prices] {
	user, err := observer.Await(userPromise)
	if err != nil && !userPromise.IsDegradable() {
		panic(err)
	}

	products, err := observer.Await(productsPromise)
	if err != nil && !productsPromise.IsDegradable() {
		panic(err)
	}

	return observer.Async(ctx,
		func() (entities.Prices, error) {
			return getPrices(user, products)
		},
	)
}
