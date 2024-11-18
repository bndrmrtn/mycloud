# Build the application
FROM golang:1.22.7-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/cloud

# Run the tests
FROM build AS run-test
RUN go test -v ./...

# Run the binary
FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app /app

CMD ["./bin/cloud", "-"]
