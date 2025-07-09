package parts

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
	"reactive-framework/internal/services"
)

func UserAndLocation(ctx context.Context) (*observer.Promise[entities.User], *observer.Promise[entities.Location]) {
	user := services.AsyncGetUser(ctx)
	location := services.AsyncGetLocation(ctx)

	return user, location
}
