package models

import (
	"database/sql"

	"github.com/gofrs/uuid"

	pb "close-spots-service/protocols/spots"
)

// Spot is a db model of a spot in db
type Spot struct {
	UUID        uuid.UUID
	Name        sql.NullString
	Website     sql.NullString
	Coordinates sql.NullString
	Description sql.NullString
	Rating      sql.NullFloat64
}

// LocationCoord is a location model with coordinates
type LocationCoord struct {
	Longitude float64
	Latitude  float64
}

// QueryRequest is a model with location and distance
type QueryRequest struct {
	LocationCoord
	Distance int
}

func ListToProto(list []*Spot) []*pb.Spot {
	spots := make([]*pb.Spot, 0, len(list))
	for _, spot := range list {
		spots = append(spots, &pb.Spot{
			Uuid:        spot.UUID.String(),
			Name:        spot.Name.String,
			Website:     spot.Website.String,
			Coordinate:  spot.Coordinates.String,
			Description: spot.Description.String,
			Rating:      float32(spot.Rating.Float64),
		})
	}

	return spots
}
