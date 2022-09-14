package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/caching"
	"github.com/idiidk/xedule-proxy/routes"
	"github.com/idiidk/xedule-proxy/xedule"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := caching.InitCacheManager()
	if err != nil {
		panic(err)
	}

	authServerEndpoint, exists := os.LookupEnv("AUTH_SERVER_ENDPOINT")
	if !exists {
		panic("AUTH_SERVER_ENDPOINT set up incorrectly!!")
	}

	authServerSecret, exists := os.LookupEnv("AUTH_SERVER_SECRET")
	if !exists {
		panic("AUTH_SERVER_SECRET set up incorrectly!!")
	}

	err = xedule.InitXedule(xedule.AuthServerConfig{Endpoint: authServerEndpoint, Secret: authServerSecret})
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{})
	routes.InitRoutes(router)

	addr, exists := os.LookupEnv("LISTEN_ADDRESS")
	if !exists {
		addr = ":8080"
	}

	router.Run(addr)
}
