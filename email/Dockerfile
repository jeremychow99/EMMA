FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go mod tidy
RUN go build -o /go-docker-email

CMD ["/go-docker-email"]