package greeting

type Greeter struct {
	greeting string
}

func New() Greeter {
	return Greeter{
		greeting: "Hello world!",
	}
}

func (g Greeter) Greet() string {
	return g.greeting
}
