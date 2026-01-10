FROM --platform=$BUILDPLATFORM golang:1.24.5-alpine AS builder

# These arguments are automatically provided by Docker buildx
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata
RUN adduser -D -g '' -u 10001 ratioapp

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Generate code with sqlc
RUN sqlc generate

# Build for the target platform
# TARGETOS and TARGETARCH are set automatically by buildx
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s" \
    -a \
    -installsuffix cgo \
    -o ratioapp ./cmd/ratioarr

FROM scratch

LABEL org.opencontainers.image.source = "https://github.com/msterhuj/ratioarr"

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/ratioapp .

USER ratioapp:ratioapp
EXPOSE 8000
ENTRYPOINT ["/app/ratioapp"]
