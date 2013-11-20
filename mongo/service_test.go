package mongo_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pjvds/counter"
	"github.com/pjvds/counter/mongo"
	"os"
	"strconv"
)

var _ = Describe("Mongo Service", func() {
	var service counter.CountService

	BeforeEach(func() {
		var err error
		var host string
		var port int

		if hostFromEnv := os.Getenv("WERCKER_MONGODB_HOST"); len(hostFromEnv) > 0 {
			host = hostFromEnv
		}

		if portFromEnv := os.Getenv("WERCKER_MONGODB_PORT"); len(portFromEnv) > 0 {
			parsed, err := strconv.Atoi(portFromEnv)
			if err != nil {
				panic(err)
			}

			port = parsed
		}

		service, err = mongo.NewCountService(fmt.Sprintf("mongodb://%v:%v/counter_test", host, port))

		Expect(err).ToNot(HaveOccured())
		Expect(service).ToNot(BeNil())
	})

	Context("when increasing a counter", func() {
		var name counter.Name
		var err error
		var beforeIncrease int
		var afterIncrease int

		BeforeEach(func() {
			name = counter.Name("my-counter")

			beforeIncrease, _ = service.Get(name)
			err = service.Increase(name)
			afterIncrease, _ = service.Get(name)
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccured())
		})

		It("should have increased by one", func() {
			Expect(afterIncrease - beforeIncrease).To(Equal(1))
		})
	})

	Context("when get a counter value of a non existing counter", func() {
		var name counter.Name
		var value int
		var err error

		BeforeEach(func() {
			name = counter.Name("not-there")
			value, err = service.Get(name)
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccured())
		})

		It("value should be 0", func() {
			Expect(value).To(Equal(0))
		})
	})
})
