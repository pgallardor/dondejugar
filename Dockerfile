FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/ -o /dondejugar-api

EXPOSE 8080

CMD ["/dondejugar-api"]