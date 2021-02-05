package main

import (
	"log"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes/builder"
)

func main() {
	routesList := builder.GenerateRoutePattern()
	log.Println(routesList)
}
