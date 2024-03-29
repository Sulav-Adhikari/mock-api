package main

import "mock-coupon-api/database"


type OfferResponse struct {
	Result bool             `json:"result"`
	Offers []database.Offer `json:"offers"`
}
