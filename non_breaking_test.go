package greeting_test

import (
	"github.com/jamesjoshuahill/greeting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greeter - setter method", func() {
	It("has a custom greeting", func() {
		greeter := greeting.New()
		greeter.Greeting("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter - alternative constructor", func() {
	It("has a custom greeting", func() {
		greeter := greeting.NewCustom("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter - builder", func() {
	It("has the default greeting", func() {
		greeter := greeting.NewBuilder().Build()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.NewBuilder().
			WithGreeting("Welcome!").
			Build()

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter - factory", func() {
	It("has a default greeting", func() {
		greeter := greeting.NewFactory("").New()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.NewFactory("Welcome!").New()

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter - modifier func", func() {
	It("has a custom greeting", func() {
		greeter := greeting.New()
		greeter = greeting.CustomGreeting(greeter, "Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("Greeter8 - constructor option param", func() {
	It("greets everyone", func() {
		greeter := greeting.New8()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New8(greeting.WithGreeting("Welcome!"))

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})
