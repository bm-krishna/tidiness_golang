package routes

import (
	"errors"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes/builder"
)

func Service(relativePath string) ([]string, error) {
	routesList, err := builder.GenerateRoutePattern(relativePath)
	if err != nil {
		return nil, errors.New("Failed to Fetch Routers Config" + err.Error())
	}
	return routesList, nil
}
