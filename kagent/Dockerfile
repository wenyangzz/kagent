# Build the kagent binary
FROM golang:1.18 as builder

# Copy in the go src
WORKDIR /go/src/github.com/kagent
COPY go.mod  go.mod
COPY go.sum  go.sum

RUN  go mod download

COPY cmd/    cmd/
COPY internal/    internal/
# Build
ENV GOPROXY "https://goproxy.cn,direct"
RUN if [ "$(uname -m)" = "ppc64le" ]; then \
        CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -a -o apiserver ./cmd/apiserver; \
    elif [ "$(uname -m)" = "aarch64" ]; then \
        CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o apiserver ./cmd/apiserver; \
    else \
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o apiserver ./cmd/apiserver; \
    fi

# Copy the kagent into a thin image
#FROM gcr.io/distroless/static:nonroot
FROM golang:1.17
WORKDIR /
COPY third_party/ third_party/
COPY --from=builder /go/src/github.com/kagent/apiserver .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL 99
ENV GRPC_GO_LOG_SEVERITY_LEVEL info
EXPOSE 9000
ENTRYPOINT ["/apiserver"]
