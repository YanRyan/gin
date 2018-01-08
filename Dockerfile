FROM golang:latest
MAINTAINER wzy gin
ADD . /go/src/gin
WORKDIR /go/src/gin
RUN go build -o server main.go
RUN chmod +x server
EXPOSE 10012
CMD ["./server"]