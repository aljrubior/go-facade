package serverGroup

import "github.com/aljrubior/go-facade/clients/hybrid/responses/common"

type Response struct {
	Id          int             `json:"id"`
	TimeCreated int64           `json:"timeCreated"`
	TimeUpdated int64           `json:"timeUpdated"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Status      string          `json:"status"`
	Servers     []common.Server `json:"servers"`
}
