package requests

import (
	"fmt"
	"github.com/aljrubior/go-facade/clients"
	"github.com/aljrubior/go-facade/config"
	"net/http"
)

func NewGetApplicationsRequest(
	config *config.CloudhubConfigClient,
	bearerToken,
	orgId,
	envId string) *GetApplicationsRequest {

	return &GetApplicationsRequest{
		config:      config,
		bearerToken: bearerToken,
		orgId:       orgId,
		envId:       envId,
	}
}

type GetApplicationsRequest struct {
	clients.BaseHttpRequest
	config *config.CloudhubConfigClient
	bearerToken,
	orgId,
	envId string
}

func (t *GetApplicationsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.ApplicationsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *GetApplicationsRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	t.AddAuthorizationHeader(req, t.bearerToken)

	return req
}
