FROM golang:1.24-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

# Copy source
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

EXPOSE 8080
ENV PORT=8080
ENV DB_PATH=/app/data/app.db

# Start the server
CMD ["./server"]
