package greeting

//////////////////////////////////////////////////
// 1. Public field
//
type Greeter1 struct {
	Greeting string
}

func New1() Greeter1 {
	return Greeter1{
		Greeting: "Hello world!",
	}
}

func (g Greeter1) Greet() string {
	return g.Greeting
}

//////////////////////////////////////////////////
// 2. Constructor field param
//
type Greeter2 struct {
	greeting string
}

func New2(greeting string) Greeter2 {
	if greeting == "" {
		greeting = "Hello world!"
	}

	return Greeter2{
		greeting: greeting,
	}
}

func (g Greeter2) Greet() string {
	return g.greeting
}

//////////////////////////////////////////////////
// 3. Constructor config struct param
//
type Greeter3 struct {
	greeting string
}

type Config struct {
	Greeting string
}

func New3(c Config) Greeter3 {
	greeting := "Hello world!"
	if c.Greeting != "" {
		greeting = c.Greeting
	}

	return Greeter3{
		greeting: greeting,
	}
}

func (g Greeter3) Greet() string {
	return g.greeting
}
