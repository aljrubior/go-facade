package cloudhub

import (
	"github.com/aljrubior/go-facade/clients/cloudhub/requests"
	"net/http"
	"time"
)

func (t DefaultHttpClient) DeleteSingleAlert(token, orgId, envId, alertId string) error {

	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewDeleteSingleAlertRequest(&t.config, token, orgId, envId, alertId).Build()

	resp, err := httpClient.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return t.ThrowError(resp)
	}

	return nil
}
