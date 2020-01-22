package apiserver

// Populate routes
func (s *Server) Routes() {
	s.v1PodRoutes()
}

func (s *Server) v1PodRoutes() {
	s.router.HandleFunc("/v1/pods", makeGetDeployedPodsEndpoint(s.log, s.pods))
}
