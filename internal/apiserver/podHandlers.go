package apiserver

import "net/http"

func makeGetDeployedPodsEndpoint(l log, repo podCollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Info("Pod list requested")
		w.Write([]byte("hello!"))
	}
}
