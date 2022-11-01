FROM golang:1.19-alpine AS BUILD

WORKDIR /app
COPY . /app

RUN go mod tidy
RUN go build -o app

#------------------
FROM alpine:3.14

WORKDIR /root
COPY --from=BUILD /app .

CMD ["./app"]