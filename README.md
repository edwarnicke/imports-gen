imports-gen is a small Go code generator designed to 'prime' a docker container in terms of both
source code download and binary caching.

## Use

Simply put a gen.go file in any package in your module containing a go:generate directive:

pkg/imports/gen.go:
```go
//go:generate go get github.com/edwarnicke/imports-gen/cmd/imports-gen
//go:generate imports-gen
```

Run:

```bash
go generate ./pkg/imports
```

This will result in the creation of a file imports.go, which imports every package imported anywhere in the module.

A docker container can then be primed with:

```dockerfile
COPY go.mod go.sum ./
COPY ./pkg/imports/ ./pkg/imports/
RUN go build ./pkg/imports/
```

The layer from the 'RUN' line will have all of the dependencies, both source and binary
for the go module.  The rest of the source code can then be copied in and built:

```dockerfile
COPY go.mod go.sum ./
COPY ./pkg/imports/ ./pkg/imports/
RUN go build ./pkg/imports/ # This line is cached
COPY . .
RUN go install ./...
```
