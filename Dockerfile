FROM --platform=$BUILDPLATFORM node:24-alpine AS tailwind

WORKDIR /app

COPY package*.json .
RUN npm install

COPY app.css input.css
RUN npx --yes @tailwindcss/cli -i input.css -o app.css --minify

FROM --platform=$BUILDPLATFORM golang:1.24.5-alpine AS builder

# These arguments are automatically provided by Docker buildx
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata gcc musl-dev sqlite-dev

# Install sqlc and templ
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Generate code with sqlc
RUN sqlc generate && \
    templ generate

# Build for the target platform
# TARGETOS and TARGETARCH are set automatically by buildx
RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s" \
    -a \
    -installsuffix cgo \
    -o ratioapp ./cmd/ratioarr

FROM alpine:latest

LABEL org.opencontainers.image.source="https://github.com/msterhuj/ratioarr"

RUN apk --no-cache add ca-certificates tzdata sqlite


WORKDIR /app

COPY --from=builder /app/ratioapp .
COPY --from=tailwind /app/app.css ./internal/static/app.css

# TODO: Add a non-root user to run the application

EXPOSE 8080
ENTRYPOINT ["/app/ratioapp"]
