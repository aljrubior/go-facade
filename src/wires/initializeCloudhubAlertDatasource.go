//go:build wireinject
// +build wireinject

package wires

import (
	cloudhubHttpClient "github.com/aljrubior/go-facade/clients/cloudhub"
	"github.com/aljrubior/go-facade/config"
	"github.com/aljrubior/go-facade/datasources/alerts"
	cloudhubAlertDatasource "github.com/aljrubior/go-facade/datasources/alerts/cloudhub"
	cloudhubService "github.com/aljrubior/go-facade/services/cloudhub"
	"github.com/google/wire"
)

func InitializeCloudhubAlertDatasource(configClient config.CloudhubConfigClient) (alerts.Datasource, error) {

	wire.Build(
		cloudhubHttpClient.NewDefaultHttpClient,
		cloudhubService.NewDefaultService,
		cloudhubAlertDatasource.NewDefaultDatasource,
		wire.Bind(new(cloudhubHttpClient.HttpClient), new(cloudhubHttpClient.DefaultHttpClient)),
		wire.Bind(new(cloudhubService.Service), new(cloudhubService.DefaultService)),
		wire.Bind(new(alerts.Datasource), new(cloudhubAlertDatasource.DefaultDatasource)),
	)

	return cloudhubAlertDatasource.DefaultDatasource{}, nil
}
