package services

import (
	"context"
	"errors"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
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
	productsPromise *promise.Promise[entities.Products],
) *promise.Promise[entities.Labels] {
	return promise.Async(ctx,
		func() (entities.Labels, error) {
			products, err := promise.Await(productsPromise)
			if err != nil && !productsPromise.IsDegradable() {
				panic(err)
			}

			return getLabels(products)
		},
	)
}

func AsyncGetLabelsWithError(
	ctx context.Context,
	productsPromise *promise.Promise[entities.Products],
) *promise.Promise[entities.Labels] {
	return promise.Async(ctx,
		func() (entities.Labels, error) {
			products, err := promise.Await(productsPromise)
			if err != nil && !productsPromise.IsDegradable() {
				panic(err)
			}

			return getLabelsWithError(products)
		},
	)
}
