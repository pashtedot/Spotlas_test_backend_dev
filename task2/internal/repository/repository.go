package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"close-spots-service/internal/metrics"
	"close-spots-service/internal/models"
)

// Repository is PostgreSQL database instance
// that implement IRepository interface
type Repository struct {
	connection *sql.DB
	metric     metrics.Metrics
}

// NewRepository create new repository instance
func NewRepository(connection *sql.DB, metric metrics.Metrics) *Repository {
	return &Repository{connection: connection, metric: metric}
}

// ListLocationsByRadius  returns a list of locations within a radius of the user's location
func (d *Repository) ListLocationsByRadius(ctx context.Context, queryRequest *models.QueryRequest) (list []*models.Spot, err error) {
	query := `SELECT
    id,
    name,
    website,
    description,
    coordinates,
    rating
FROM
    "MY_TABLE"
WHERE
    ST_DWithin(
            coordinates::geography,
            ST_Point($1,$2)::geography,
            $3
        )
ORDER BY
  CASE
    WHEN ST_Distance(coordinates::geography, ST_Point($1,$2)::geography) < 50 THEN rating
    ELSE ST_Distance(coordinates::geography, ST_Point($1,$2)::geography)
  END;
`

	return d.queryToSpots(query, queryRequest)
}

// ListLocationsInSquareByHalvedSide returns a list of locations within a square with a side equal to two distances
func (d *Repository) ListLocationsInSquareByHalvedSide(ctx context.Context, queryRequest *models.QueryRequest) (list []*models.Spot, err error) {
	query := `SELECT
    id,
    name,
    website,
    description,
    coordinates,
    rating
FROM
    "MY_TABLE"
WHERE
    coordinates && ST_Envelope(ST_Buffer(ST_Point($1,$2)::geography, $3)::geometry)
ORDER BY
  CASE
    WHEN ST_Distance(coordinates::geography, ST_Point($1,$2)::geography) < 50 THEN rating
    ELSE ST_Distance(coordinates::geography, ST_Point($1,$2)::geography)
  END;
`

	return d.queryToSpots(query, queryRequest)
}

func (d *Repository) queryToSpots(query string, queryRequest *models.QueryRequest) (list []*models.Spot, err error) {
	rows, err := d.connection.Query(query, queryRequest.Longitude, queryRequest.Latitude, queryRequest.Distance)
	if err != nil {
		return nil, fmt.Errorf("cant list loc(%v,%v) by radius %v: %w", queryRequest.Longitude, queryRequest.Latitude, queryRequest.Distance, err)
	}
	defer rows.Close()

	for rows.Next() {
		spot := &models.Spot{}
		err = rows.Scan(&spot.UUID, &spot.Name, &spot.Website, &spot.Description, &spot.Coordinates, &spot.Rating)
		if err != nil {
			return nil, fmt.Errorf("cant scan list loc(%v,%v) by radius %v: %w", queryRequest.Longitude, queryRequest.Latitude, queryRequest.Distance, err)
		}
		list = append(list, spot)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during the iteration over the rows: %w", err)
	}

	return
}
