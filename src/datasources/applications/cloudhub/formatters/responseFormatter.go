package formatters

import "github.com/aljrubior/go-facade/clients/cloudhub/responses"

func NewResponseFormatter(response *[]responses.ApplicationResponse) ResponseFormatter {

	return ResponseFormatter{
		response: response,
	}
}

type ResponseFormatter struct {
	response *[]responses.ApplicationResponse
}
