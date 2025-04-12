# syntax=docker/dockerfile:1

# -----------------------------------------------------------------------------
# Build Stage
# -----------------------------------------------------------------------------
  ARG GO_VERSION=1.22.0
  FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
  
  WORKDIR /src
  
  # Copy go mod files first to leverage Docker cache
  COPY go.mod go.sum ./
  RUN go mod download
  
  # Copy entire source code (backend code + frontend /out)
  COPY . .
  
  # Build Go backend
  ARG TARGETARCH
  RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -tags docker -o /bin/server ./cmd
  
  # -----------------------------------------------------------------------------
  # Final Stage
  # -----------------------------------------------------------------------------
  FROM alpine:latest AS final
  
  # Install runtime dependencies
  RUN apk --no-cache add ca-certificates tzdata
  
  # Create non-root user
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
  
  WORKDIR /
  
  # Copy server binary from build stage
  COPY --from=build /bin/server /bin/server
  
  # ✅ Copy static frontend files from build stage
  COPY --from=build /out /out
  
  EXPOSE 8080
  
  # Run backend
  ENTRYPOINT ["/bin/server"]
  