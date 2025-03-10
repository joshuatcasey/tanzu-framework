# Copyright 2022 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# Build from publicly reachable source by default, but allow people to re-build images on
# top of their own trusted images.
ARG BUILDER_BASE_IMAGE=golang:1.18

# Build the manager binary
FROM $BUILDER_BASE_IMAGE as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY apis/run/ apis/run/
COPY util/ util/

COPY tkr/go.mod tkr/go.mod
COPY tkr/go.sum tkr/go.sum

WORKDIR /workspace/tkr
RUN  go mod download
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer

# Copy the go source
COPY tkr/ ./

# Build
ARG LD_FLAGS
ENV LD_FLAGS="$LD_FLAGS "'-extldflags "-static"'
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags "$LD_FLAGS" -o manager controller/tkr-source/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/tkr/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
