FROM --platform=$BUILDPLATFORM node:24-alpine AS tailwind

WORKDIR /app

COPY package*.json .
RUN npm install

# Copy necessary files for Tailwind CSS generation
COPY app.css tailwind.config.js ./
COPY internal/views ./internal/views
RUN npx --yes @tailwindcss/cli -i app.css -o output.css --minify

FROM --platform=$BUILDPLATFORM golang:1.24.5-alpine AS builder

# These arguments are automatically provided by Docker buildx
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

# Install sqlc and templ
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
# Copy CSS from tailwind stage to make it available for Go embed
COPY --from=tailwind /app/output.css ./internal/static/app.css
# Generate code with sqlc
RUN sqlc generate && \
    templ generate

# Build for the target platform
# TARGETOS and TARGETARCH are set automatically by buildx
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s" \
    -o ratioapp ./cmd/ratioarr

FROM alpine:latest

LABEL org.opencontainers.image.source="https://github.com/msterhuj/ratioarr"

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/ratioapp .

# TODO: Add a non-root user to run the application

EXPOSE 8080
ENTRYPOINT ["/app/ratioapp"]
