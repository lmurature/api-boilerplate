FROM golang:1.20.2

ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}
ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV AUTH_KEY=${AUTH_KEY}
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -a -o app ./cmd/api

EXPOSE 8080

CMD ["./app"]