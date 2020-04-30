# Learning GO

A series of experiments, scripts and findings from learning Go.

## Technology

- [Go](https://golang.org/)

## Resources

- [Go Playground](https://play.golang.org/)
- [The little Go Book](https://www.openmymind.net/assets/go/go.pdf)

## Why Go?

- Simple Language
- Simple standard library
- Sits nicely between low level system applications & higher level applications
- Add more....

## Language Key Points

### Go is a compiled

- Compilation is very quick.
- Compiled languages run faster interpreted languages.
- Executables can be run without additional dependencies.

### Go is statically typed

- Variable must be of declared with a specific type

### Go has garbage collection built in

- saves time clearing up after yourself

### Go is not an OO Language

- It doesn't support objects
- It doesn't support inheritance
- It doesn't support polymorphism
- It doesn't support overloading

## Docs

- [Official Documentation](https://golang.org)
- [Packages](https://golang.org/pkg)

Note: Go docs are available offline.

```
godoc -http=:6060

docs available at: http:localhost:6060
```

## Development

To run a script:

```bash
$ cd src

$ go run {filename}.go

```

To build a script:

```bash
$ cd src

$ go build {filename}.go

// Run newly built executable
$ ./{filename}.go

```
