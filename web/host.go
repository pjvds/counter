package web

import (
	"fmt"
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

	name := counter.Name(r.URL.Path)

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
	w.Write([]byte(fmt.Sprintf("{'%v': %v}", name, value)))
}
