package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
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
	userPromise *observer.Promise[entities.User],
	locationPromise *observer.Promise[entities.Location],
) *observer.Promise[entities.Products] {
	user, err := observer.Await(userPromise)
	if err != nil && !userPromise.IsDegradable() {
		panic(err)
	}

	location, err := observer.Await(locationPromise)
	if err != nil && !locationPromise.IsDegradable() {
		panic(err)
	}

	return observer.Async(ctx,
		func() (entities.Products, error) {
			return getProducts(user, location)
		},
	)
}
