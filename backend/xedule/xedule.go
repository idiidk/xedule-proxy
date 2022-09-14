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

func InitXedule(authServerConfig AuthServerConfig) error {
	authEndpoint = authServerConfig.Endpoint
	authSecret = authServerConfig.Secret

	res, err := GetAuthServerResponse()
	if err != nil {
		return err
	}
	xeduleEndpoint = res.Config.XeduleURL

	return nil
}

func queryXedule(endpoint string, object any) error {
	err := caching.UnmarshalCache(endpoint, &object)
	if err == nil {
		return nil
	}

	client, err := GetAuthenticatedHttpClient()
	if err != nil {
		return err
	}

	res, err := client.Get(fmt.Sprintf("%s/api/%s", xeduleEndpoint, endpoint))
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&object)
	if err != nil {
		return err
	}

	caching.MarshalCache(endpoint, object)
	return nil
}

func GetGroups() (*[]models.XeduleGroup, error) {
	var groups *[]models.XeduleGroup = new([]models.XeduleGroup)
	err := queryXedule("group", groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func GetOrganisationalUnits() (*[]models.XeduleOrganisationalUnit, error) {
	var organisationalUnits *[]models.XeduleOrganisationalUnit = new([]models.XeduleOrganisationalUnit)
	err := queryXedule("organisationalUnit", organisationalUnits)
	if err != nil {
		return nil, err
	}

	return organisationalUnits, nil
}
