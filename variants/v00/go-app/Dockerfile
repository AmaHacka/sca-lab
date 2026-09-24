FROM golang:1.19-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/configconv .

FROM alpine:3.14
COPY --from=build /out/configconv /usr/local/bin/configconv
EXPOSE 8080
CMD ["configconv"]
