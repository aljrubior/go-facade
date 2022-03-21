package hybrid

import (
	"github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/alert/requests"
)

func (t DefaultDatasource) UpdateAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error) {

	return t.hybridService.UpdateSingleAlert(token, orgId, envId, alertId, request)
}
