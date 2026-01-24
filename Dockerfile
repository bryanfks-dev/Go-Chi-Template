ARG GO_VERSION=1.25.5
ARG MAIN_GO_FILE=cmd/http/main.go
ARG APP_ENVIRONMENT=production

ARG OS=alpine
ARG USERNAME=golang

FROM golang:${GO_VERSION}-${OS} AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the project
COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
    go build -o main -ldflags="-s -w" ${MAIN_GO_FILE}

FROM ${OS}:latest AS runtime

# Create a non-root user to run the application
RUN addgroup --gid 1001 --system ${USERNAME} && \
    adduser --system --uid 1001 --ingroup ${USERNAME} appuser

WORKDIR /app

ENV APP_ENVIRONMENT=${APP_ENVIRONMENT}

COPY --from=builder --chown=appuser:${USERNAME} /app/main .

CMD ["./main"]
