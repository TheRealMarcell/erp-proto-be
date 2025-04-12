# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.0
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
ARG TARGETARCH
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -tags docker -o /bin/server ./cmd

FROM alpine:latest AS final

RUN apk --update add \
    ca-certificates \
    tzdata \
    && \
    update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=build /bin/server /bin/

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
