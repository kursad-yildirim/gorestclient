FROM docker.io/golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o moviecollectionclient

FROM docker.io/alpine:latest
COPY --from=build /app/moviecollectionclient /app/moviecollectionclient
WORKDIR /app
EXPOSE 8080
CMD ["./moviecollectionclient"]
