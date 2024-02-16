FROM golang:1.21.5-alpine as build

WORKDIR /app

COPY go.mod .

RUN go mod download

#ENV GOOS linux
#ENV GOARCH amd64

COPY . .

RUN go build -o app

#FROM --platform=linux/amd64 alpine
FROM alpine

WORKDIR /app

COPY --from=build /app/app /app

ENTRYPOINT [ "/app/app" ]