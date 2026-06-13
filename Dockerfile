FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD [ "./app" ]