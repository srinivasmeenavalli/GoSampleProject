# alcl-go-function-template

This is a template for golang based knative functions

## HOWTOs

### Build

```bash
➜  alcl-go-function-template git:(readme) ✗ make build
CGO_ENABLED=0 go build -o ./bin/alcl-go-function-template
````

### Test

```bash
➜  alcl-go-function-template git:(readme) ✗ make test
CGO_ENABLED=0 go test -v ./...
?       alcl-go-function-template      [no test files]
```

### All together

```bash
➜  alcl-go-function-template git:(readme) ✗ make 
CGO_ENABLED=0 go mod download
CGO_ENABLED=0 go install github.com/golangci/golangci-lint/cmd/golangci-lint
CGO_ENABLED=0 go test -v ./...
?       alcl-go-function-template      [no test files]
CGO_ENABLED=0 golangci-lint run
CGO_ENABLED=0 go build -o ./bin/alcl-go-function-template
```

Makes sure golangci-lint is installed in the PATH
