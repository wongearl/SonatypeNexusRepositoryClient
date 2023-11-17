
FROM golang:1.20 as builder

ARG GOPROXY=https://goproxy.cn,direct
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN GOPROXY=${GOPROXY} go mod download
# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY main.go main.go
# Build
RUN make snr-client

FROM alpine:3.18.4
WORKDIR /workspace
RUN apk add sudo
COPY --from=builder /workspace/bin/snr-client snr-client
RUN chmod +x snr-client
