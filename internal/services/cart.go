package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
)

func getCart(
	user entities.User,
	products entities.Products,
	prices entities.Prices,
	labels entities.Labels,
) (entities.Cart, error) {
	items := make([]entities.CartItem, len(products))

	for _, product := range products {
		item := entities.CartItem{
			ID:          product.ID,
			Quantity:    product.Quantity,
			Description: product.Description,
		}

		if price, ok := prices[product.ID]; ok {
			item.Price = price
		}

		if label, ok := labels[product.ID]; ok {
			item.Label = label.Value
		}

		items = append(items, item)
	}

	return entities.Cart{
		UserID: user.Id,
		Items:  items,
	}, nil
}

func AsyncGetCart(
	ctx context.Context,
	userPromise *promise.Promise[entities.User],
	productsPromise *promise.Promise[entities.Products],
	pricesPromise *promise.Promise[entities.Prices],
	labelsPromise *promise.Promise[entities.Labels],
) *promise.Promise[entities.Cart] {
	return promise.Async(ctx,
		func() (entities.Cart, error) {
			user, err := promise.Await(userPromise)
			if err != nil && !userPromise.IsDegradable() {
				panic(err)
			}

			products, err := promise.Await(productsPromise)
			if err != nil && !productsPromise.IsDegradable() {
				panic(err)
			}

			prices, err := promise.Await(pricesPromise)
			if err != nil && !pricesPromise.IsDegradable() {
				panic(err)
			}

			labels, err := promise.Await(labelsPromise)
			if err != nil && !labelsPromise.IsDegradable() {
				panic(err)
			}

			return getCart(user, products, prices, labels)
		},
	)
}
