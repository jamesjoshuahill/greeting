package greeting_test

import (
	"github.com/jamesjoshuahill/greeting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greeter", func() {
	It("greets everyone", func() {
		greeter := greeting.New()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})
})
