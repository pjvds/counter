package mongo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pjvds/counter"
	"github.com/pjvds/counter/mongo"
)

var _ = Describe("Mongo Service", func() {
	var service counter.CountService

	BeforeEach(func() {
		var err error
		service, err = mongo.NewCountService("mongodb://localhost/counter_test")

		Expect(err).ToNot(HaveOccured())
		Expect(service).ToNot(BeNil())
	})

	Context("when increasing a counter that does not exist", func() {
		var name counter.Name
		var err error

		BeforeEach(func() {
			name = counter.Name("my-counter")
			err = service.Increase(name)
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccured())
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
