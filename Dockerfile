FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /src/
COPY go.mod go.sum ./
COPY app ./app

RUN go mod tidy

# COPY . /.
RUN go build -C ./app -o main

FROM alpine:latest

WORKDIR /root/
COPY --from=build /src/app/main .
EXPOSE 8080

CMD ["./main"]
