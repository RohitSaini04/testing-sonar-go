FROM golang:1.22.2-alpine3.19 AS build

# COPY go.sum go.mod /.
WORKDIR /src/
COPY . .
RUN go mod tidy

# COPY . /.
RUN go build -C app -o main

# FROM alpine:latest

# WORKDIR /root/
# COPY --from=build /src/app/main .
# EXPOSE 8080

# CMD ["./main"]
