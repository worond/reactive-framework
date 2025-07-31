package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
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
	userPromise *promise.Promise[entities.User],
	productsPromise *promise.Promise[entities.Products],
) *promise.Promise[entities.Prices] {
	return promise.Async(ctx,
		func() (entities.Prices, error) {
			user, err := promise.Await(userPromise)
			if err != nil && !userPromise.IsDegradable() {
				panic(err)
			}

			products, err := promise.Await(productsPromise)
			if err != nil && !productsPromise.IsDegradable() {
				panic(err)
			}

			return getPrices(user, products)
		},
	)
}
