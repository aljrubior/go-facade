package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/serverGroup"
	model2 "github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newServerGroupDetails(serverGroup serverGroup.Response) model2.Details {

	return model2.Details{
		Servers: t.formatCommonServers(&serverGroup.Servers),
	}
}
