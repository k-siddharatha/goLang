FROM golang:latest
RUN mkdir -p /go/src/app/
COPY . /go/src/app/
WORKDIR /go/src/app/
EXPOSE 8080
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["main"]
