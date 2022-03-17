package alert

import (
	"encoding/json"
	"github.com/aljrubior/amc-ui-rest-facade/controllers/headers"
	"log"
	"net/http"
)

func (t DefaultController) GetAlerts(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(headers.AuthorizationHttpHeader)
	orgId := r.Header.Get(headers.OrganizationIdHttpHeader)
	envId := r.Header.Get(headers.EnvironmentIdHttpHeader)

	resp, err := t.alertService.GetAlerts(token, orgId, envId)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(resp)
}
