package alert

import (
	"errors"
	"fmt"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
)

func (t DefaultService) GetAlertHistory(token, orgId, envId, product, alertId string) (*[]alerts.AlertHistoryResponse, error) {

	datasource, ok := t.datasources[product]

	if !ok {
		// TODO: Implement this
		return nil, errors.New(fmt.Sprintf("Alerts not supported for product '%s'", product))
	}

	return datasource.GetAlertHistory(token, orgId, envId, alertId)
}
