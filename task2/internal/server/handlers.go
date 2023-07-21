package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "close-spots-service/protocols/spots"

	"close-spots-service/internal/models"
)

// ISpotsService is interface for Application Service Layer
type ISpotsService interface {
	// ListSpots return list of spots near point
	ListSpots(ctx context.Context, longitude, latitude float64, dist int, searchType models.QueryType) (list []*models.Spot, err error)
	Close()
}

// Handlers is GRPC handlers implementation
// for Spots Protocol
type Handlers struct {
	srv ISpotsService

	*pb.UnimplementedCloseSpotsServiceServer
}

// NewHandlers create new WhitelistHandlers instance
func NewHandlers(srv ISpotsService) *Handlers {
	return &Handlers{srv: srv}
}

func (h *Handlers) ListSpotsNearPoint(ctx context.Context, req *pb.ListSpotsNearPointRequest) (*pb.ListSpotsNearPointResponse, error) {
	spots, err := h.srv.ListSpots(ctx, req.GetLongitude(), req.GetLatitude(), int(req.GetLength()), models.QueryTypeToDomain(req.GetSearchType()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cant ListSpots: %s", err.Error())
	}

	return &pb.ListSpotsNearPointResponse{
		Spots: models.ListToProto(spots),
		Total: int64(len(spots)),
	}, nil
}
