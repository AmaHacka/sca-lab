FROM golang:1.21.0-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/reports .

FROM alpine:3.17
COPY --from=build /out/reports /usr/local/bin/reports
EXPOSE 8080
CMD ["reports"]
