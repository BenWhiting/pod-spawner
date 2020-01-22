package apiserver

import "net/http"

func makeGetDeployedPodsEndpoint(l log, repo podCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Info("requesting pods")
		w.WriteHeader(http.StatusNotImplemented)
	}
}
