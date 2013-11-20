package web_test

import (
	"encoding/json"
	"fmt"
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
	var service *countertest.CountService

	BeforeEach(func() {
		service = countertest.NewCountService()
		host = NewServiceHost(service)
	})

	Context("Handling correct timer request", func() {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "http://localhost/my-timer", nil)

		BeforeEach(func() {
			service.GetFunc = func(name counter.Name) (int, error) {
				return 15, nil
			}
			host.ServeHTTP(response, request)
		})

		AfterEach(func() {
			service.GetFunc = nil
		})

		It("Should set Content-Type header", func() {
			Expect(response.HeaderMap).Should(HaveKey("Content-Type"))
			Expect(response.HeaderMap.Get("Content-Type")).Should(Equal("application/json"))
		})

		It("Should return a body", func() {
			Expect(response.Body.Len()).Should(BeNumerically(">", 0))
		})

		It("Should return value returned by service", func() {
			var body struct {
				Value int `json:"value"`
			}
			decoder := json.NewDecoder(response.Body)
			if err := decoder.Decode(&body); err != nil {
				Fail(fmt.Sprintf("Error unmarshalling body: %v, %v", string(response.Body.Bytes()), err))
			}

			Expect(body.Value).Should(Equal(15))
		})
	})
})
