package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
	"time"
)

func getProducts(user entities.User, location entities.Location) (entities.Products, error) {
	time.Sleep(100 * time.Millisecond)

	return entities.Products{
		1111: {ID: 1111, Quantity: 1},
		2222: {ID: 2222, Quantity: 2},
		3333: {ID: 3333, Quantity: 3},
	}, nil
}

func AsyncGetProducts(
	ctx context.Context,
	userPromise *promise.Promise[entities.User],
	locationPromise *promise.Promise[entities.Location],
) *promise.Promise[entities.Products] {
	return promise.Async(ctx,
		func() (entities.Products, error) {
			user, err := promise.Await(userPromise)
			if err != nil && !userPromise.IsDegradable() {
				panic(err)
			}

			location, err := promise.Await(locationPromise)
			if err != nil && !locationPromise.IsDegradable() {
				panic(err)
			}

			return getProducts(user, location)
		},
	)
}
