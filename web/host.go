package web

import (
	"encoding/json"
	"github.com/pjvds/counter"
	"net/http"
)

type ServiceHost struct {
	service counter.CountService
}

func NewServiceHost(service counter.CountService) *ServiceHost {
	return &ServiceHost{
		service: service,
	}
}

func (host *ServiceHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := counter.Name(r.URL.Path[1:])

	if err := host.service.Increase(name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	value, err := host.service.Get(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(struct {
		Name  counter.Name `json:"name"`
		Value int          `json:"value"`
	}{
		Name:  name,
		Value: value,
	})
}
