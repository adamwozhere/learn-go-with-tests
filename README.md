# Learn Go With Tests

This repo contains my working files for [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests).

## Test Driven Development

Test Driven Development follows a cycle:

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

## Installing Go

Install Go from [go.dev/doc/install](https://go.dev/doc/install).
You can check it is installed with the command `go version`.
Modules are used by go to manage dependencies (similar to NPM), the module name is usually named the same as the module path.
To create a project:

```bash
mkdir my-project
cd my-project
go mod init <modulepath>
```

## Linting

Linting can be done with VSCode: Install the offical Go extention and select `golangci-lint` as the formatter - it will need to install, but it is better than the default linter.
`go fmt` can be used to format Go code. But it is easier to just configure the VSCode extention to use `golanci-lint` to format on save.

## Go doc

Installed go packages have documentation which can be accessed with `godoc`.
In newer versions of Go, you will need to manually install using `go install golang.org/x/tools/cmd/godoc@latest`.
To use `godoc` use `godoc -http=localhost:8000`, if you then visit `localhost:8000/pkg`, you will then see all the packages installed on your system.
