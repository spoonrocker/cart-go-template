FROM ubuntu:18.04

RUN apt-get update && apt-get -y upgrade

RUN apt-get install -y curl && apt-get install -y wget && apt-get install -y sudo

COPY src /go/src/github.com/Kleiber/cart-go-template/src/

COPY go.mod /go/src/github.com/Kleiber/cart-go-template/go.mod

COPY go.sum /go/src/github.com/Kleiber/cart-go-template/go.sum

ADD setup.sh /

RUN chmod +x /setup.sh

RUN /setup.sh

WORKDIR /go/src/github.com/Kleiber/cart-go-template/src/

EXPOSE 3000

CMD ["./app"]

