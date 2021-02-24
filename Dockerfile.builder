
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

RUN mkdir /app 
WORKDIR /app/ 

COPY ./go.mod .
COPY ./phoneNumberServer.go .

RUN mkdir logger
COPY ./logger/ ./logger/

RUN mkdir webService
COPY ./webService/ ./webService/

RUN go build -o phoneNumberServer 

EXPOSE 8080

ENTRYPOINT  /app/phoneNumberServer
