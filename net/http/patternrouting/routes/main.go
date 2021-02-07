package routes

import (
	"errors"
	"log"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes/builder"
)

func Service(relativePath string) ([]string, error) {
	log.Println(relativePath, "Routes")
	routesList, err := builder.GenerateRoutePattern(relativePath)
	if err != nil {
		return nil, errors.New("Failed to Fetch Routers Config" + err.Error())
	}
	log.Println(routesList)
	return routesList, nil
}
