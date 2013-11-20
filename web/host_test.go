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
		timerName := "my-timer"

		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "http://localhost/"+timerName, nil)
		timerValue := 15

		BeforeEach(func() {
			service.GetFunc = func(name counter.Name) (int, error) {
				return timerValue, nil
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

		Context("The body", func() {
			var body struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			}

			BeforeEach(func() {
				decoder := json.NewDecoder(response.Body)
				if err := decoder.Decode(&body); err != nil {
					Fail(fmt.Sprintf("Error unmarshalling body: %v, %v", string(response.Body.Bytes()), err))
				}
			})

			It("Should return name as provided by request", func() {
				Expect(body.Name).Should(Equal(timerName))
			})

			It("Should return value as provided by service", func() {
				Expect(body.Value).Should(Equal(timerValue))
			})
		})
	})
})
