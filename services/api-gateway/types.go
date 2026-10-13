package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID string `json:"user_id"`
	PickupCoordinate types.Coordinate `json:"pickup_coordinate"`
	DestinationCoordinate types.Coordinate `json:"destination_coordinate"`
}