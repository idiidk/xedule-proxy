package xedule

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/idiidk/xedule-proxy/caching"
	"github.com/idiidk/xedule-proxy/xedule/models"
)

func GetAuthServerResponse() (*models.AuthServerResponse, error) {
	cacheKey := "authServerResponse"

	var authServerResponse models.AuthServerResponse
	err := caching.UnmarshalCache(cacheKey, &authServerResponse)

	if err != nil {
		// no cache
		res, err := http.Get(fmt.Sprintf("%s?secret=%s", authEndpoint, authSecret))
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&authServerResponse)
		if err != nil {
			return nil, err
		}

		caching.MarshalCache(cacheKey, authServerResponse)
		return &authServerResponse, nil
	} else {
		// cached
		return &authServerResponse, nil
	}
}

func GetAuthenticatedHttpClient() (*http.Client, error) {
	authServerResponse, err := GetAuthServerResponse()
	if err != nil {
		return nil, err
	}

	jar, err := getCookieJar(*authServerResponse)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Jar: jar,
	}, nil
}

func getCookieJar(response models.AuthServerResponse) (*cookiejar.Jar, error) {
	cookieJar, _ := cookiejar.New(nil)

	var cookies []*http.Cookie
	var finalDomain string

	for _, cookie := range response.Cookies {
		finalDomain = cookie.Domain

		cookies = append(cookies, &http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}

	domain, err := url.Parse(fmt.Sprintf("https://%s", finalDomain))
	if err != nil {
		return nil, err
	}

	cookieJar.SetCookies(domain, cookies)
	return cookieJar, nil
}
