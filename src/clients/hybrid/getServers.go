package hybrid

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/clients/hybrid/requests"
	"github.com/aljrubior/go-facade/clients/hybrid/responses/server"
	"io/ioutil"
	"net/http"
	"time"
)

func (t DefaultHttpClient) GetServers(token, orgId, envId string) (*server.DataResponse, error) {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewGetServerGroupsRequest(&t.config, token, orgId, envId).Build()

	resp, err := httpClient.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, t.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response server.DataResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
