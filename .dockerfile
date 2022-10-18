FROM golang:1.19
RUN mkdir /build
ADD go.mod go.sum app.go /build/
WORKDIR /build
RUN go build