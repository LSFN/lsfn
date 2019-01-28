[![Go Report Card](https://goreportcard.com/badge/github.com/LSFN/lsfn)](https://goreportcard.com/report/github.com/LSFN/lsfn)

# lsfn
Lower Shields, Fire Nukes

## Installing & Building

1. [Install Go](https://developers.google.com/protocol-buffers/docs/downloads). Minimum version 1.11.
2. Install `protoc` onto your `PATH`. See the [Protocol Buffers download instructions](https://developers.google.com/protocol-buffers/docs/downloads) or follow your platform's installation instructions (e.g., `apt-get install protobuf-compiler`).
3. Install Open Dynamics. Either by following the [build instructions](https://bitbucket.org/odedevs/ode/) or by following your platform's instructions (e.g., `apt-get install libode-dev`).
4. Clone this repository. You can develop inside your `GOPATH` with `go get -d github.com/LSFN/lsfn` and setting the environment variable `GO111MODULE=on`, or outside of it with a normal `git clone` of this repository.
5. Build the `cmd/environment` and `cmd/ship` binaries. If inside the GOPATH use `go build github.com/LSFN/lsfn/cmd/environment`, otherwise enter the correct directory and execute `go build` there.

## Developing

* When making changes to the `.proto` files under `api/proto`, run `make generate` to invoke `protoc` to recreate the `.pb.go` files. These files should be commited with the changes to the `.proto` files.
