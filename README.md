# Golang Merit Money

Web application based on the concept of the Merit Money reward scheme (https://management30.com/blog/merit-money-a-reward-scheme-that-works/). Uses Gin as Web framework and Go template rendering for the HTML rendering.

Integrates dependabot and Github Workflows for binary releases.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/qaware/golang-merit-money)

## Running the application

Install Go from here: https://golang.org/doc/install there are usually packages for it on every system.

Run it with `go run main.go`.

The server runs by default on port 8080 so you can view it in the browser under localhost:8080.

## Running tests

```
go test -v ./...
```

### Measure code coverage

```
go test -coverprofile=coverage.out -coverpkg=./... ./...
go tool cover -func coverage.out
```

## Github CI Pipeline

The workflows do the following:
- Runs golangci-lint for code quality
- Runs tests and measures code coverage which is sent to Codecov.
- On tags, it builds a new Github Release with binaries for all popular operating systems and generates a changelog based on conventional commits

Definitions are found under `.github/workflows`

## Pre-commit

A pre-commit config is inside the Repo. It runs golangci-lint, format and gomod tidy and unit tests.

First install pre-commit with:

`pip install pre-commit`

Then install the commit hooks by running inside the repository folder:

`pre-commit install`

Now the commit hooks are automatically run on every git commit.

Optionally you can run `pre-commit run --all-files` to see the output without a commit.