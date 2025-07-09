package services

import (
	"context"
	"reactive-framework/internal/entities"
	"reactive-framework/internal/observer"
	"time"
)

func getLocation() (entities.Location, error) {
	time.Sleep(100 * time.Millisecond)

	return entities.Location{
		Lat: 42.1,
		Lon: 32.4,
	}, nil
}

func AsyncGetLocation(ctx context.Context) *observer.Promise[entities.Location] {
	return observer.Async(ctx, getLocation)
}
