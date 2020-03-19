package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/kubernetes_scheduling_simulation/pkg/api"
)

func main() {
	r := mux.NewRouter()

	api_url := r.PathPrefix("/api/v1").Subrouter()

	api_url.HandleFunc("/nodes", api.Nodes).Methods(http.MethodGet)
	api_url.HandleFunc("/pods", api.Pods).Methods(http.MethodGet)
	api_url.HandleFunc("/namespaces/{namespace}/pods/{pod}/binding", api.Binding).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8001", r))
}
