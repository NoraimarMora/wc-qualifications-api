FROM golang:1.20-bullseye as builder

ENV GO111MODULE=on

# Create directories
#RUN mkdir -p /go/src/wc-qualifications-api
#WORKDIR /go/src/wc-qualifications-api

WORKDIR /app

# Copy app and run go mod.
COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd/wc-qualifications-api
#ADD . /go/src/wc-qualifications-api

# Copy and run the app.
#RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/wc-qualifications-api

#FROM golang:1.20-bullseye
#WORKDIR /root/
#COPY --from=0 /go/src/wc-qualifications-api .
#RUN mkdir -p /root/downloads
#RUN mkdir -p /root/files
#VOLUME /root/downloads
#VOLUME /root/files
ENTRYPOINT ["./app"]
