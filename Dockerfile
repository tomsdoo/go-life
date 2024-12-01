FROM golang:1.22

COPY ./ /usr/local/
WORKDIR /usr/local/app

RUN go build -o lifegame main.go

CMD ["./lifegame"]
