FROM golang:1.19-alpine

# Install GCC and build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /tradehub ./cmd/server

EXPOSE 8000

CMD ["/tradehub"]

