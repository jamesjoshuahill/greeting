package greeting

//////////////////////////////////////////////////
// 4. Setter method
//
func (g *Greeter) Greeting(greeting string) {
	g.greeting = greeting
}

//////////////////////////////////////////////////
// 5. a) Alternative constructor
//
func NewCustom(greeting string) Greeter {
	return Greeter{
		greeting: greeting,
	}
}

//////////////////////////////////////////////////
// 5. b) Constructor param with default value
//       (not possible in Go)
//
//func New(greeting = "Hello world!" string) Greeter2 {
//	return Greeter2{
//		greeting: greeting,
//	}
//}

//////////////////////////////////////////////////
// 5. c) Overload constructor
//       (not possible in Go)
//

//////////////////////////////////////////////////
// 6. Builder
//
type Builder struct {
	greeting string
}

func NewBuilder() *Builder {
	return &Builder{
		greeting: "Hello world!",
	}
}

func (b *Builder) WithGreeting(g string) *Builder {
	b.greeting = g
	return b
}

func (b *Builder) Build() Greeter {
	return Greeter{
		greeting: b.greeting,
	}
}

//////////////////////////////////////////////////
// 6. Factory
//
type Factory struct {
	greeting string
}

func NewFactory(greeting string) Factory {
	if greeting == "" {
		greeting = "Hello world!"
	}

	return Factory{
		greeting: greeting,
	}
}

func (f Factory) New() Greeter {
	return Greeter{
		greeting: f.greeting,
	}
}

//////////////////////////////////////////////////
// 7. Modifier func
//
func CustomGreeting(g Greeter, greeting string) Greeter {
	g.greeting = greeting
	return g
}

//////////////////////////////////////////////////
// 8. Constructor option param
//
type Greeter8 struct {
	greeting string
}

type Option func(*Greeter8)

func WithGreeting(greeting string) Option {
	return func(g *Greeter8) {
		g.greeting = greeting
	}
}

func New8(options ...Option) Greeter8 {
	g := Greeter8{
		greeting: "Hello world!",
	}

	for _, option := range options {
		option(&g)
	}

	return g
}

func (g Greeter8) Greet() string {
	return g.greeting
}
