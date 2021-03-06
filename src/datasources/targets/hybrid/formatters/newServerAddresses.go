package formatters

import (
	"github.com/aljrubior/go-facade/clients/hybrid/responses/common"
	"github.com/aljrubior/go-facade/datasources/targets/hybrid/model"
)

func (t ResponseFormatter) newServerAddresses(server common.Server) []model.Address {

	var addresses []model.Address

	for _, v := range server.Addresses {
		addresses = append(addresses, model.Address{
			Ip:               v.Ip,
			NetworkInterface: v.NetworkInterface,
		})
	}

	return addresses
}
