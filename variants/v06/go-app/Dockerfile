FROM golang:1.19-bullseye AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/sshorch .

FROM ubuntu:18.04
COPY --from=build /out/sshorch /usr/local/bin/sshorch
EXPOSE 8080
CMD ["sshorch"]
