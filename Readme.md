# Golang for Beginners

This repository is designed to help beginners get started with Golang. Below is a quick guide to installing, setting up, and using Go effectively.

## Prerequisites

Ensure that you have Go installed. To check the installed version of Go, run:

```bash
go version
```

## Setting Up the Environment

Configure your Go environment:

```bash
go env -w GO111MODULE=off
```

Run your Go application:

- To run a Go application from the current directory:

```bash
go run .
```

- To run a specific Go file:

```bash
go run filename.go
```

For help with Go commands:

```bash
go help
```

## Package Management

Go code is grouped into **packages**, and packages are grouped into **modules**. To explore packages, visit [pkg.go.dev](https://pkg.go.dev).

Initialize your Go module using:

```bash
go mod init [module-path]
go mod init hello
```

### Adding Packages

To add external packages to your Go module:

```bash
go get -v rsc.io/quote
```

### Managing Dependencies

The `go mod tidy` command will ensure that your dependencies are in order:

- It **adds** missing dependencies.
- It **removes** unused ones.

Run it with:

```bash
go mod tidy
```

## Building Go Projects

Compile and build your Go project:

```bash
go build
```

- To build a specific Go file:

```bash
go build filename.go
```

## Listing Go Modules

To list the modules and their dependencies:

```bash
go list
```

## Additional Resources

- Official Go Documentation: [golang.org/doc](https://golang.org/doc)
- Explore Go Packages: [pkg.go.dev](https://pkg.go.dev)
