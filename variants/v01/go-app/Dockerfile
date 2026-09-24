FROM golang:1.18-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/notes .

FROM alpine:3.13
COPY --from=build /out/notes /usr/local/bin/notes
EXPOSE 8080
CMD ["notes"]
