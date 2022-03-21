package hybrid

import "github.com/aljrubior/amc-ui-rest-facade/clients/responses/alerts"

func (t DefaultDatasource) GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	return t.hybridService.GetAlertHistory(token, orgId, envId, alertId)
}
