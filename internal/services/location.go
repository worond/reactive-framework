package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/promise"
	"time"
)

func getLocation() (entities.Location, error) {
	time.Sleep(200 * time.Millisecond)

	return entities.Location{
		Lat: 42.1,
		Lon: 32.4,
	}, nil
}

func AsyncGetLocation(ctx context.Context) *promise.Promise[entities.Location] {
	return promise.Async(ctx, getLocation)
}
