package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/luci/go-render/render"
	v1 "k8s.io/api/core/v1"
)

func Pods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//TODO : put this somewhere accessible by everybody
	// Mock responses files
	box := packr.New("mocks", "../../mock")
	response, err := box.FindString("get-pods.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("get pods")

	//TODO : handle more cases
	query := r.URL.Query()
	fieldSelector := query.Get("fieldSelector")
	params := strings.Split(fieldSelector, "=")

	// TODO : handle more cases, also have a more robust test
	if params[0] == "spec.nodeName" && params[1] == "" {
		w.Write([]byte(string(response)))
	}
}

func Nodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("get nodes")

	box := packr.New("mocks", "../../mock")
	response, err := box.FindString("get-nodes.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//// TODO : return a fixed mock response
	w.Write([]byte(response))
}

func Binding(w http.ResponseWriter, r *http.Request) {
	// Only prints what it receives, so we can check what messages are passed
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	log.Println("Got post request")

	// Query parameters
	if namespace, ok := pathParams["namespace"]; ok {
		log.Printf("namespace : %s\n", namespace)
	}
	if pod, ok := pathParams["pod"]; ok {
		log.Printf("pod : %s\n", pod)
	}

	// Body parameters
	var p *v1.Binding
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println(render.Render(p))
	w.WriteHeader(http.StatusCreated)

	box := packr.New("mocks", "../../mock")
	response, err := box.FindString("response_success.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(response))
}
