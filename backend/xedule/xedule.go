package xedule

import (
	"encoding/json"
	"fmt"

	"github.com/idiidk/xedule-proxy/caching"
	"github.com/idiidk/xedule-proxy/xedule/models"
)

var authEndpoint string
var authSecret string

var xeduleEndpoint string

func InitXedule(authServerConfig models.AuthServerConfig) error {
	authEndpoint = authServerConfig.Endpoint
	authSecret = authServerConfig.Secret

	res, err := GetAuthServerResponse()
	if err != nil {
		return err
	}
	xeduleEndpoint = res.Config.XeduleURL

	return nil
}

func GetGroups() ([]models.XeduleGroup, error) {
	cacheKey := "groups"

	var groups []models.XeduleGroup
	err := caching.UnmarshalCache(cacheKey, &groups)
	if err == nil {
		return groups, nil
	}

	client, err := GetAuthenticatedHttpClient()
	if err != nil {
		return nil, err
	}

	res, err := client.Get(fmt.Sprintf("%s/api/group", xeduleEndpoint))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&groups)
	if err != nil {
		return nil, err
	}

	caching.MarshalCache(cacheKey, groups)

	return groups, nil
}
