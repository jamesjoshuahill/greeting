package greeting_test

import (
	"github.com/jamesjoshuahill/greeting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("1. Constructor field param", func() {
	It("greets everyone", func() {
		greeter := greeting.New1("")

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New1("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("2. Constructor config struct param", func() {
	It("greets everyone", func() {
		greeter := greeting.New2(greeting.Config{})

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New2(greeting.Config{Greeting: "Welcome!"})

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})
