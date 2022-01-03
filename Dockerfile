FROM golang:latest as build

ADD . /app
WORKDIR /app

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/realtime.exe .

EXPOSE 8080

ENTRYPOINT ["/app/realtime.exe"]
