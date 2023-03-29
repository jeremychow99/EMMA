FROM golang:1.20-alpine
ENV TZ=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go mod tidy
RUN go build -o /go-docker-automaintenance

CMD ["/go-docker-automaintenance"]