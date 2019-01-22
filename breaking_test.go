package greeting_test

import (
	"github.com/jamesjoshuahill/greeting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greeter1 - public field", func() {
	It("greets everyone", func() {
		greeter := greeting.New1()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New1()
		greeter.Greeting = "Welcome!"

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter2 - constructor field param", func() {
	It("greets everyone", func() {
		greeter := greeting.New2("")

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New2("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter3 - constructor config struct param", func() {
	It("greets everyone", func() {
		greeter := greeting.New3(greeting.Config{})

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New3(greeting.Config{Greeting: "Welcome!"})

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})
