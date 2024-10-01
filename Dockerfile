FROM golang:1.20-bullseye as builder

ENV GO111MODULE=on

WORKDIR /app

# Copy app and run go mod.
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd/wc-qualifications-api

ENTRYPOINT ["./app"]
