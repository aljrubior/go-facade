package transformers

import (
	"github.com/aljrubior/go-facade/clients/cloudhub/responses"
	"github.com/aljrubior/go-facade/datasources/targets/cloudhub/model"
	"github.com/aljrubior/go-facade/model/responses/target"
)

func NewDefaultTransformer(applications *[]responses.ApplicationResponse) *DefaultTransformer {
	return &DefaultTransformer{
		applications: applications,
	}
}

type DefaultTransformer struct {
	applications *[]responses.ApplicationResponse
}

func (t DefaultTransformer) Transform() (*[]target.Response, error) {

	result := make([]target.Response, 0)

	if len(*t.applications) > 0 {
		result = append(result, model.NewTarget())
	}

	return &result, nil

}
