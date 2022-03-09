package services

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/clients/cloudhub"
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alertController/requests"
)

func NewDefaultCloudhubService(httpClient cloudhub.HttpClient) DefaultCloudhubService {

	return DefaultCloudhubService{
		httpClient: httpClient,
	}

}

type DefaultCloudhubService struct {
	httpClient cloudhub.HttpClient
}

func (t DefaultCloudhubService) GetAlerts(token string) (*[]alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetAlerts(token)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultCloudhubService) GetSingleAlert(token, orgId, envId, alertId string) (*[]alerts.AlertResponse, error) {

	resp, err := t.httpClient.GetSingleAlert(token, orgId, envId, alertId)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (t DefaultCloudhubService) PostAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.AlertResponse, error) {

	body, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.PostAlert(token, orgId, envId, body)

	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
