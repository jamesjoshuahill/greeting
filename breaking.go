package greeting

//////////////////////////////////////////////////
// 1. Constructor field param
//
type Greeter1 struct {
	greeting string
}

func New1(greeting string) Greeter1 {
	if greeting == "" {
		greeting = "Hello world!"
	}

	return Greeter1{
		greeting: greeting,
	}
}

func (g Greeter1) Greet() string {
	return g.greeting
}

//////////////////////////////////////////////////
// 2. Constructor config struct param
//
type Greeter2 struct {
	greeting string
}

type Config struct {
	Greeting string
}

func New2(c Config) Greeter2 {
	greeting := "Hello world!"
	if c.Greeting != "" {
		greeting = c.Greeting
	}

	return Greeter2{
		greeting: greeting,
	}
}

func (g Greeter2) Greet() string {
	return g.greeting
}
