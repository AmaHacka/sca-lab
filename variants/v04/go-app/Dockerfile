FROM golang:1.21.0-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/weather .

FROM debian:bullseye-slim
COPY --from=build /out/weather /usr/local/bin/weather
EXPOSE 8080
CMD ["weather"]
