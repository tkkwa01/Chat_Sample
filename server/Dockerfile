FROM golang:1.20 as builder

RUN go install github.com/cosmtrek/air@latest

ENV CGO_ENABLED=0

ENV GOOS=linux

ENV GOARCH=amd64

ENV ROOTPATH=/go/app

WORKDIR ${ROOTPATH}

COPY go.mod go.sum .air.toml ./

RUN go mod download


EXPOSE 80

CMD ["air", "-c", ".air.toml"]