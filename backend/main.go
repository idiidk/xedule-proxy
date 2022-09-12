package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/idiidk/xedule-proxy/caching"
	"github.com/idiidk/xedule-proxy/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := caching.InitCacheManager()
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
