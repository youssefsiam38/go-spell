FROM golang:1.11

ENV GOBIN="$GOPATH/bin"
WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
