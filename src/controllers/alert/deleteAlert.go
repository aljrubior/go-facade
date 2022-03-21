package alert

import (
	"encoding/json"
	"github.com/aljrubior/go-facade/controllers/headers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (t DefaultController) DeleteAlert(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(headers.AuthorizationHttpHeader)
	orgId := r.Header.Get(headers.OrganizationIdHttpHeader)
	envId := r.Header.Get(headers.EnvironmentIdHttpHeader)

	vars := mux.Vars(r)

	productId, ok := vars["productId"]

	if !ok {
		// TODO: Implement
		http.Error(w, "productId is missing in parameters", http.StatusBadRequest)
	}

	alertId, ok := vars["alertId"]

	if !ok {
		// TODO: Implement
		http.Error(w, "alertId is missing in parameters", http.StatusBadRequest)
	}

	err := t.alertService.DeleteSingleAlert(token, orgId, envId, productId, alertId)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("")
}
