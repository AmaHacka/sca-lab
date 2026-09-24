FROM golang:1.19-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/avatars .

FROM alpine:3.15
COPY --from=build /out/avatars /usr/local/bin/avatars
EXPOSE 8080
CMD ["avatars"]
