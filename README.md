# Greeting

This is the starting point for a workshop in [Go](https://golang.org).

## Get

```bash
$ go get github.com/jamesjoshuahill/greeting
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
```

## Test

The test suite can be run using `go test`, or the `ginkgo` CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
$ go get -u github.com/onsi/gomega
$ go get -u github.com/onsi/ginkgo/ginkgo
$ ginkgo -r -p
```

## Exercise

The starting point is [`greeting.go`](https://github.com/jamesjoshuahill/greeting/blob/master/greeting.go). You have:

- a struct with a private field
- the constructor accepts no parameters
- the constructor sets a default value for the private field

Identify as many ways as possible to:

- add support for custom greetings
- continue to offer the default greeting, `Hello world!`
