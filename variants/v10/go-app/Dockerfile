FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/auth .

FROM scratch
COPY --from=build /out/auth /auth
EXPOSE 8080
ENTRYPOINT ["/auth"]
