FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/tracker .

FROM alpine:3.16
COPY --from=build /out/tracker /usr/local/bin/tracker
EXPOSE 8080
CMD ["tracker"]
