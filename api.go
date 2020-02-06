package main

import (
	"context"
	"googlemaps.github.io/maps"
)

func getRestaurants(apiKey string) ([]string, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return []string{}, err
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
		return []string{}, err
	}
	result := make([]string, len(nearby.Results))
	for i, v := range nearby.Results {
		result[i] = v.Name
	}
	return result, nil
}
