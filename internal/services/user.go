package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
	"time"
)

func getUser() (entities.User, error) {
	time.Sleep(100 * time.Millisecond)

	return entities.User{
		Id:      1,
		Session: "aaa",
	}, nil
}

func AsyncGetUser(ctx context.Context) *promise.Promise[entities.User] {
	return promise.Async(ctx, getUser)
}
