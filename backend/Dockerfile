FROM golang:1.21

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

CMD ["./backend"]

