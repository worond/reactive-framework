package services

import (
	"context"
	"errors"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
	"time"
)

func getLabels(products entities.Products) (entities.Labels, error) {
	time.Sleep(100 * time.Millisecond)

	labels := make(entities.Labels, len(products))
	for _, product := range products {
		labels[product.ID] = entities.Label{Value: "LABEL"}
	}

	return labels, nil
}

func getLabelsWithError(products entities.Products) (entities.Labels, error) {
	time.Sleep(100 * time.Millisecond)

	return nil, errors.New("LABELS ERROR")
}

func AsyncGetLabels(
	ctx context.Context,
	productsPromise *observer.Promise[entities.Products],
) *observer.Promise[entities.Labels] {
	products, err := observer.Await(productsPromise)
	if err != nil && !productsPromise.IsDegradable() {
		panic(err)
	}

	return observer.Async(ctx,
		func() (entities.Labels, error) {
			return getLabels(products)
		},
	)
}

func AsyncGetLabelsWithError(
	ctx context.Context,
	productsPromise *observer.Promise[entities.Products],
) *observer.Promise[entities.Labels] {
	products, err := observer.Await(productsPromise)
	if err != nil && !productsPromise.IsDegradable() {
		panic(err)
	}

	return observer.Async(ctx,
		func() (entities.Labels, error) {
			return getLabelsWithError(products)
		},
	)
}
