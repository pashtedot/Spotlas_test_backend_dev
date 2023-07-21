package models

import (
	pb "close-spots-service/protocols/spots"
)

// QueryType Domain type
type QueryType int

const (
	InvalidType QueryType = iota
	Circle
	Square
)

func QueryTypeToDomain(p pb.QueryType) QueryType {
	switch p {
	case pb.QueryType_QUERY_TYPE_CIRCLE:
		return Circle
	case pb.QueryType_QUERY_TYPE_SQUARE:
		return Square
	default:
		return InvalidType
	}
}
func (qt QueryType) String() string {
	switch qt {
	case Circle:
		return "CircleType"
	case Square:
		return "SquareType"
	default:
		return "InvalidType"
	}
}
