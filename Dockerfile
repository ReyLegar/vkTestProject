FROM golang:1.22.1

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
RUN go build -o vkTestProject ./cmd/vkTestProject/main.go

CMD ["./vkTestProject"]