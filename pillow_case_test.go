package main_test

import (
	"io/ioutil"
	"net/http"
	pillow "github.com/njbennett/pillow-case"
	sheets "google.golang.org/api/sheets/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/njbennett/pillow-case/pillow-casefakes"
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
		XIt("returns the values in that sheet", func() {
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

var _ = Describe("Handler", func() {
	Context("when given a response writer and a request", func() {
		It("writes the request path to the writer", func() {
			fakeWriter := new(FakeResponseWriter)
			request, _ := http.NewRequest("GET", "/string", nil)
			pillow.Handler(fakeWriter, request)
			Expect(fakeWriter.WriteCallCount()).To(Equal(1))
			Expect(fakeWriter.WriteArgsForCall(0)).To(Equal([]byte("string")))
		})
	})
})

var _ = Describe("LookupHandler", func() {
	Context("when the request path is /A1", func() {
		It("writes 'hello'", func() {
			fakeWriter := new(FakeResponseWriter)
			request, _ := http.NewRequest("GET", "/A1", nil)
			pillow := new(pillow.Pillow)
			pillow.LookupHandler(fakeWriter, request)
			Expect(fakeWriter.WriteCallCount()).To(Equal(1))
			Expect(fakeWriter.WriteArgsForCall(0)).To(Equal([]byte("hello")))
		})
	})
		
	Context("when the request path is /B2", func() {
		It("writes 'famicom'", func() {
			fakeWriter := new(FakeResponseWriter)
			request, _ := http.NewRequest("GET", "/B2", nil)
			pillow := new(pillow.Pillow)
			pillow.LookupHandler(fakeWriter, request)
			Expect(fakeWriter.WriteCallCount()).To(Equal(1))
			Expect(fakeWriter.WriteArgsForCall(0)).To(Equal([]byte("famicom")))
		})
	})
})
