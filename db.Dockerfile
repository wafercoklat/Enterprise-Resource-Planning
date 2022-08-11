FROM golang:latest

LABEL maintener = "Riza"

WORKDIR /app

COPY go.mod . 

COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build 

CMD ["./REVAMP-PHP-GO"] 