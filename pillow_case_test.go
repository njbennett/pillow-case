package main_test

import (
	pillow "github.com/njbennett/pillow-case"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Write", func() {
	It("accepts a string", func() {
		err := pillow.Write("test")
		Expect(err).ToNot(HaveOccurred())
	})
})
