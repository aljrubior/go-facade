package hybrid

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses"
	"github.com/aljrubior/go-facade/clients/hybrid/responses/cluster"
	"github.com/aljrubior/go-facade/clients/hybrid/responses/server"
	"github.com/aljrubior/go-facade/clients/hybrid/responses/serverGroup"
	"github.com/aljrubior/go-facade/clients/responses/alerts"
	"github.com/aljrubior/go-facade/controllers/alert/requests"
)

type Service interface {
	GetAlerts(token, orgId, envId string) (*[]alerts.Response, error)
	GetSingleAlert(token, orgId, envId, alertId string) (*alerts.Response, error)
	UpdateSingleAlert(token, orgId, envId, alertId string, request requests.AlertRequest) (*[]alerts.Response, error)
	CreateAlert(token, orgId, envId string, request requests.AlertRequest) (*[]alerts.Response, error)
	DeleteSingleAlert(token, orgId, envId, alertId string) error
	GetAlertHistory(token, orgId, envId, alertId string) (*[]alerts.AlertHistoryResponse, error)
	GetResourceAlertHistory(token, orgId, envId, resourceId string) (*[]alerts.ResourceAlertHistory, error)

	GetApplications(token, orgId, envId string) (*[]responses.ApplicationResponse, error)
	GetServers(token, orgId, envId string) (*[]server.Response, error)
	GetServerGroups(token, orgId, envId string) (*[]serverGroup.Response, error)
	GetClusters(token, orgId, envId string) (*[]cluster.Response, error)
}
