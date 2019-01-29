package greeting_test

import (
	"github.com/jamesjoshuahill/greeting"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("3. Make field public", func() {
	It("greets everyone", func() {
		greeter := greeting.New3()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New3()
		greeter.Greeting = "Welcome!"

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("4. Setter method", func() {
	It("has a custom greeting", func() {
		greeter := greeting.New()
		greeter.Greeting("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("5. Alternative constructor", func() {
	It("has a custom greeting", func() {
		greeter := greeting.NewCustom("Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("6. Builder", func() {
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

var _ = Describe("7. Factory", func() {
	It("has a default greeting", func() {
		greeter := greeting.NewFactory("").New()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.NewFactory("Welcome!").New()

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("8. Modifier func", func() {
	It("has a custom greeting", func() {
		greeter := greeting.New()
		greeter = greeting.CustomGreeting(greeter, "Welcome!")

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})

var _ = Describe("9. Constructor option param", func() {
	It("greets everyone", func() {
		greeter := greeting.New9()

		Expect(greeter.Greet()).To(Equal("Hello world!"))
	})

	It("has a custom greeting", func() {
		greeter := greeting.New9(greeting.WithGreeting("Welcome!"))

		Expect(greeter.Greet()).To(Equal("Welcome!"))
	})
})
