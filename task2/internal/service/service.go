package service

import (
	"context"
	"fmt"

	"close-spots-service/internal/metrics"
	"close-spots-service/internal/models"
)

type IRepository interface {
	ListLocationsByRadius(ctx context.Context, queryRequest *models.QueryRequest) (list []*models.Spot, err error)
	ListLocationsInSquareByHalvedSide(ctx context.Context, queryRequest *models.QueryRequest) (list []*models.Spot, err error)
}

// SpotsService is implementation for service Layer
type SpotsService struct {
	repository  IRepository
	metric      metrics.Metrics
	stopService chan struct{}
}

// NewSpotsService create new SpotsService instance
func NewSpotsService(
	metric metrics.Metrics,
	repository IRepository,
) (*SpotsService, error) {
	return &SpotsService{
		repository:  repository,
		metric:      metric,
		stopService: make(chan struct{}),
	}, nil
}

// ListSpots return all spots
func (p *SpotsService) ListSpots(ctx context.Context, longitude, latitude float64, dist int, searchType models.QueryType) (list []*models.Spot, err error) {
	request := &models.QueryRequest{
		LocationCoord: models.LocationCoord{
			Longitude: longitude,
			Latitude:  latitude,
		},
		Distance: dist,
	}

	switch searchType {
	case models.Circle:
		list, err = p.repository.ListLocationsByRadius(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("cant get  Locations By Radius: %w", err)
		}

		return

	case models.Square:
		list, err = p.repository.ListLocationsInSquareByHalvedSide(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("cant get  Locations in square: %w", err)
		}

		return

	default:
		return nil, fmt.Errorf("searchType %s is not supported", searchType.String())
	}
}

func (p *SpotsService) Close() {
	close(p.stopService)
}
