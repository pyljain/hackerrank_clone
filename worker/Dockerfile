FROM golang:1.14 as builder
WORKDIR /go/src/github.com/patnaikshekhar/worker

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app 

FROM ubuntu:latest

RUN apt-get update \
  && apt-get install -y python3-pip python3-dev \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && pip3 install --upgrade pip

RUN apt-get install curl gnupg -yq \
    && curl -sL https://deb.nodesource.com/setup_12.x | bash \
    && apt-get install nodejs -yq

CMD ["./app"]
RUN mkdir /problem
COPY --from=builder /app .