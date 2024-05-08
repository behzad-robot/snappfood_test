FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o snappfood cmd/web/main.go
EXPOSE 5400
CMD ["./snappfood"]