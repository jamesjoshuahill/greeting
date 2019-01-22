# Greeting

| Time | Duration | Description                             |
|:-----|:---------|:----------------------------------------|
| h:00 |          | Start                                   |
| h:05 | 5m       | Introduction                            |
| h:10 | 5m       | Breaking patterns                       |
| h:15 | 15m      | Non-breaking patterns                   |
| h:30 | 10m      | Share non-breaking patterns             |
| h:40 | 15m      | Mob on this scenario in production code |
| h:55 | 5m       | Wrap up                                 |
| h:00 |          | End                                     |

## Introduction

Let's explore patterns in [Go](https://golang.org).
This workshop focuses on structs and constructors.

The starting point is [`greeting.go`](https://github.com/jamesjoshuahill/greeting/blob/master/greeting.go).
You have:

- a struct with a private field
- the constructor accepts no parameters
- the constructor sets a default value for the private field

**Goal:** Find as many ways as possible to support custom greetings. For example: `"Welcome!"`

## Breaking patterns

The API of a package in Go includes the constants, variables and types that are public.

This includes public structs and their public fields and methods.

A few ways to add support for custom greetings by breaking the API are shown in
[`breaking.go`](https://github.com/jamesjoshuahill/greeting/blob/master/breaking.go).

## Non-breaking patterns

Let's do a breadth exercise to find as many ways to extend the behaviour of the
`Greeting` struct without breaking the package API.

Please note:

- **Please write all your ideas in Go.**
- The default greeting should be `"Hello world!"`.
- Dirty code is fine. It only needs to compile.
- If you think of something possible in another language, please make a note.
- Please try not to filter out your ideas. Try as many as you can. Go wide!

### Get

```bash
$ go get github.com/jamesjoshuahill/greeting
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
```

### Test

The test suite can be run using `go test`, or the [`ginkgo`](https://onsi.github.io/ginkgo/) CLI:

```bash
$ cd $GOPATH/src/github.com/jamesjoshuahill/greeting
$ go get -u github.com/onsi/gomega
$ go get -u github.com/onsi/ginkgo/ginkgo
$ ginkgo
```
