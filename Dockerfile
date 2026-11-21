# Usage:
#   docker run --cap-add CAP_SYS_MODULE -v /lib/modules:/lib/modules:ro wadey/synology-tun

FROM --platform=$BUILDPLATFORM golang:1-alpine AS build

ARG TARGETOS=linux
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} \
    go build -trimpath -ldflags="-s -w" -o /synology-tun .

FROM scratch

COPY --from=build /synology-tun /synology-tun

ENTRYPOINT ["/synology-tun"]
