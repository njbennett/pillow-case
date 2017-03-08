package main_test

import (
	"net/http"

	pillow "github.com/njbennett/pillow-case"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("hello world handler", func() {
	It("Returns 'Hello world'", func() {
		server := ghttp.NewServer()
		server.AppendHandlers(pillow.HelloWorldHandler)
		request, err := http.NewRequest("Get", server.URL(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(request).To(Equal("Hello world"))
	})
})
