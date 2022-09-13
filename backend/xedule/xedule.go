package xedule

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/idiidk/xedule-proxy/caching"
	"github.com/idiidk/xedule-proxy/xedule/models"
)

func AuthServerCookieToJar(authServerCookies []models.AuthServerCookie) (*cookiejar.Jar, error) {
	var cookieDomain string
	var cookies []*http.Cookie

	for _, cookie := range authServerCookies {
		cookieDomain = cookie.Domain
		cookieToSet := http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		}

		cookies = append(cookies, &cookieToSet)
	}

	domainUrl, err := url.Parse(fmt.Sprintf("https://%s", cookieDomain))
	if err != nil {
		return nil, err
	}

	cookieJar.SetCookies(domainUrl, cookies)
	return cookieJar, nil
}

func GetAuthenticatedCookieJar(endpoint string, secret string) (*cookiejar.Jar, error) {
	cacheKey := "cookie-jar"
	var authServerCookies []models.AuthServerCookie

	err := caching.UnmarshalCache(context.Background(), cacheKey, cookieJar)
	if err == nil {
		fmt.Println("WE ARE USING CACHE LETSGO EZ PZ LEMON SQUEEZE")
		return cookieJar, nil
	}

	var authServerCookies []models.AuthServerCookie
	resp, err := http.Get(fmt.Sprintf("%s?secret=%s", endpoint, secret))
	if err != nil {
		return nil, err
	}

	json.NewDecoder(resp.Body).Decode(&authServerCookies)

	caching.MarshalCache(context.Background(), cacheKey, cookieJar)

	return cookieJar, nil
}
