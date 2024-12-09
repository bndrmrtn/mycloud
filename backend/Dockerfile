# Build the application
FROM golang:1.22.7-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cp /app/config.yaml /app/config.file || cp /app/config.yml /app/config.file

RUN go build -o bin/mycloud

# Run the tests
FROM build AS run-test
RUN go test -v ./...

# Run the binary
FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/bin /app
COPY --from=build /app/config.file /app/config.yaml

CMD ["./mycloud", "-"]
