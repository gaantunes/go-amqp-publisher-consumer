FROM golang:1.16-alpine

WORKDIR /app

COPY . . 
RUN go mod download

RUN go build ./src/consumer/consumer.go

CMD [ "./consumer" ]