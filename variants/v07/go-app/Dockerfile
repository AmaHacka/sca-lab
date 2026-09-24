FROM golang:1.18-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/chatbot .

FROM alpine:3.12
COPY --from=build /out/chatbot /usr/local/bin/chatbot
EXPOSE 8080
CMD ["chatbot"]
