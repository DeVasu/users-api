FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

EXPOSE 9092
# EXPOSE 3306


RUN go build -o /token-jwt

CMD ["/token-jwt"]