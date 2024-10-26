FROM golang:1.22.7-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/cloud
CMD ["./bin/cloud", "-"]
