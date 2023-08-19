FROM golang:1.21-alpine as dev
RUN mkdir -p /app

COPY . /app
WORKDIR /app
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0
RUN go install github.com/swaggo/swag/cmd/swag@latest
CMD ["air", "-c", ".air.toml"]