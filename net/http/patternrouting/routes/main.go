package routes

import (
	"log"
	"os"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes/builder"
)

func main() {
	relaivePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to read relative path")
	}
	routesList := builder.GenerateRoutePattern(relaivePath)
	log.Println(routesList)
}
