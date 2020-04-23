# Learning GO

A series of experiments, scripts and findings from learning Go.

## Technology

* [Go](https://golang.org/)

## Docs

* [Packages](https://golang.org/pkg)

Note: Go docs are available offline.

```
godoc -http=:6060

docs available at: http:localhost:6060
```


## Tutorials

* [The little Go Book](https://www.openmymind.net/assets/go/go.pdf) 

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