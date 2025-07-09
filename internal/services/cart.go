package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
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

func AwaitGetCart(
	ctx context.Context,
	userPromise *observer.Promise[entities.User],
	productsPromise *observer.Promise[entities.Products],
	pricesPromise *observer.Promise[entities.Prices],
	labelsPromise *observer.Promise[entities.Labels],
) (entities.Cart, error) {
	user, err := observer.Await(userPromise)
	if err != nil && !userPromise.IsDegradable() {
		panic(err)
	}

	products, err := observer.Await(productsPromise)
	if err != nil && !productsPromise.IsDegradable() {
		panic(err)
	}

	prices, err := observer.Await(pricesPromise)
	if err != nil && !pricesPromise.IsDegradable() {
		panic(err)
	}

	labels, err := observer.Await(labelsPromise)
	if err != nil && !labelsPromise.IsDegradable() {
		panic(err)
	}

	return getCart(user, products, prices, labels)
}
