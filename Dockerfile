FROM golang:latest as build

ADD . /app
WORKDIR /app

RUN go get -u all
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/realtime.exe .

ENTRYPOINT ["/app/realtime.exe"]
