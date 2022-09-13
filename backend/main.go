package main

import (
	"net/url"
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

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies([]string{})
	routes.InitRoutes(router)

	authServerAddr, _ := os.LookupEnv("AUTH_SERVER_ADDRESS")
	authServerSecret, _ := os.LookupEnv("AUTH_SERVER_SECRET")

	_, err = xedule.GetAuthenticatedCookieJar(authServerAddr, authServerSecret)
	if err != nil {
		panic(err)
	}

	_, err = xedule.GetAuthenticatedCookieJar(authServerAddr, authServerSecret)
	if err != nil {
		panic(err)
	}

	c, err := xedule.GetAuthenticatedCookieJar(authServerAddr, authServerSecret)
	if err != nil {
		panic(err)
	}

	u, _ := url.Parse("https://sa-han.xedule.nl")
	println(c.Cookies(u))

	addr, exists := os.LookupEnv("LISTEN_ADDRESS")
	if !exists {
		addr = ":8080"
	}

	router.Run(addr)
}
