package main

import (
	"context"
	"googlemaps.github.io/maps"
)
// Restaurant is a struct representing a search result
type Restaurant struct {
	Name     string
	Location maps.LatLng
	PlaceID  string
}

func getRestaurants(apiKey string) ([]Restaurant, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return []Restaurant{}, err
	}
	r := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: 49.480690,
			Lng: 8.472023,
		},
		Language: "german",
		Radius:   500,
		Type:     maps.PlaceTypeRestaurant,
	}
	nearby, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		return []Restaurant{}, err
	}
	result := make([]Restaurant, len(nearby.Results))
	for i, v := range nearby.Results {
		result[i] = Restaurant{
			Name:     v.Name,
			Location: v.Geometry.Location,
			PlaceID:  v.PlaceID,
		}
	}
	return result, nil
}
