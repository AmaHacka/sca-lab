FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/shortener .

FROM debian:buster-slim
COPY --from=build /out/shortener /usr/local/bin/shortener
EXPOSE 8080
CMD ["shortener"]
