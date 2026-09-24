FROM golang:1.18-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/rss .

FROM debian:stretch-slim
COPY --from=build /out/rss /usr/local/bin/rss
EXPOSE 8080
CMD ["rss"]
