package library

import (
	"context"
	"googlemaps.github.io/maps"
)

// NewClient はClientを作成
func NewClient() (*maps.Client, error) {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBKVoLU4aaZ5_NZkElgOANEzybtbVRaQUY"))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetGeoInfoFromAddress(ctx context.Context, residentAddress string) (maps.GeocodingResult, error) {
	c, err := NewClient()
	var result maps.GeocodingResult
	if err != nil {
		return result, err
	}

	r := &maps.GeocodingRequest{
		Address:residentAddress,
	}
	results, err := c.Geocode(context.Background(), r)

	if err != nil {
		return result, err
	}
	result = results[0]
	return result, nil
}
