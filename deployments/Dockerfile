FROM golang:1.21

WORKDIR /opt/app

COPY . .

RUN go mod download

RUN go build -o /usr/bin/rednote-app .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/usr/bin/rednote-app"]
