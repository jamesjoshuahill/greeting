# Greeting

This is the starting point for a workshop in [Go](https://golang.org).

## Get

```bash
$ go get github.com/jamesjoshuahill/greeting
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
```

## Test

The test suite can be run using `go test`, or the [`ginkgo`](https://onsi.github.io/ginkgo/) CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
$ go get -u github.com/onsi/gomega
$ go get -u github.com/onsi/ginkgo/ginkgo
$ ginkgo -r -p
```

## Breadth exercise

The starting point is [`greeting.go`](https://github.com/jamesjoshuahill/greeting/blob/master/greeting.go). You have:

- a struct with a private field
- the constructor accepts no parameters
- the constructor sets a default value for the private field

**Goal:** Find as many ways as possible to support custom greetings. For example: `Welcome!`

Please note:
- The default greeting should be `Hello world!`.
- Please write all your ideas in Go. As long as they compile, they're valid!
- If you think of something possible in another language. Please make a note.
- Please don't filter out your ideas. Look for as many as you can.
