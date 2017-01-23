package main_test

import (
	"io/ioutil"
	pillow "github.com/njbennett/pillow-case"
	sheets "google.golang.org/api/sheets/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FetchCredentials", func() {
	Context("when a credentials file is present", func() {
		It("returns client credentials", func() {
			dat, err := ioutil.ReadFile("./credentials")
			creds, err := pillow.FetchCredentials()
			Expect(err).ToNot(HaveOccurred())
			Expect(creds).To(Equal(dat))
		})
	})
})

var _ = Describe("Read", func() {
	Context("when called on the test sheet", func() {
		It("returns the values in that sheet", func() {
			Expect(pillow.Read("A1")).To(Equal("hello"))
			Expect(pillow.Read("B2")).To(Equal("famicom"))
		})
	})
})

var _ = Describe("Write", func() {
	Context("when given a string and a sheets service", func() {
		It("returns no error", func() {
			srv := new(sheets.Service)
			err := pillow.Write("test", srv)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
