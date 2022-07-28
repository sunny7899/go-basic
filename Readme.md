Go will be the server language
https://golang.org/doc/

go version
go env -w GO111MODULE=off
go run .
go run filename.go
go help
go version

for packages info-
pkg.go.dev

go mod init hello

to add packages-
go get -v rsc.io/quote

Go code is grouped into packages, and packages are grouped into modules

Start your module using the go mod init command

go mod init [module-path]

this will find all dependencies, add the missing and remove the unused dependencies.
go mod tidy

go build

go build filename.go

go list
