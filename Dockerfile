# syntax=docker/dockerfile:1

###############################################################################
# Build Stage - Frontend
###############################################################################
FROM node:20-alpine AS frontend-build

WORKDIR /app

# Copy only frontend files
COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Build the frontend
RUN npm run build

###############################################################################
# Build Stage - Backend
###############################################################################
FROM --platform=$BUILDPLATFORM golang:1.22 AS backend-build

WORKDIR /src

# Copy Go mod files and download modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build Go backend
ARG TARGETARCH
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH go build -tags docker -o /bin/server ./cmd

###############################################################################
# Final Production Image
###############################################################################
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

# Copy Go server binary
COPY --from=backend-build /bin/server /bin/server

# ✅ Copy built frontend from frontend-build stage
COPY --from=frontend-build /out /out

# Expose port
EXPOSE 8080

# Start server
ENTRYPOINT ["/bin/server"]
