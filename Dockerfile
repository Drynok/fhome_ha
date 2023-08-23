FROM golang:1.19

RUN apt-get update
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

# Copy the project files
COPY . .

# Build the binary
RUN go build -o main *.go

COPY start.sh ./

CMD ["./start.sh"]
