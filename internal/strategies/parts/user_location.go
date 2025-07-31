package parts

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
	"reactive-framework/internal/services"
)

func UserAndLocation(ctx context.Context) (*promise.Promise[entities.User], *promise.Promise[entities.Location]) {
	user := services.AsyncGetUser(ctx)
	location := services.AsyncGetLocation(ctx)

	return user, location
}
