package main_test

import (
	pillow "github.com/njbennett/pillow-case"
	sheets "google.golang.org/api/sheets/v4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write", func() {
	Context("when given a string and a sheets service", func() {
		It("returns no error", func() {
			srv := new(sheets.Service)
			err := pillow.Write("test", srv)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
