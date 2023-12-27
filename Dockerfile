FROM golang:1.20

WORKDIR /go/src/app
COPY . .

RUN go build -o /go/bin/app main.go

CMD ["/go/bin/app"]
