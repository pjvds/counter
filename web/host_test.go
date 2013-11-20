package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pjvds/counter"
	"github.com/pjvds/counter/countertest"
	. "github.com/pjvds/counter/web"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Host", func() {
	var host *ServiceHost
	var service counter.CountService

	BeforeEach(func() {
		service = countertest.NewNullService()
		host = NewServiceHost(service)
	})

	Context("Handling correct timer request", func() {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "http://localhost/my-timer", nil)

		BeforeEach(func() {
			host.ServeHTTP(response, request)
		})

		It("Should set Content-Type header", func() {
			Expect(response.HeaderMap).Should(HaveKey("Content-Type"))
			Expect(response.HeaderMap.Get("Content-Type")).Should(Equal("application/json"))
		})
	})
})
