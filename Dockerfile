FROM golang:1.14-alpine AS BUILDER
RUN apk --no-cache add git
WORKDIR /go/src/github.com/cabogabo/cart-api/
ADD . .
RUN go mod download
RUN go mod verify
RUN go build -o ./bin/serve ./cmd/main.go

FROM alpine:3.9
WORKDIR /usr/bin
COPY --from=BUILDER /go/src/github.com/cabogabo/cart-api/bin /go/bin
EXPOSE 3000
CMD [ "/go/bin/serve" ]