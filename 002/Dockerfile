FROM golang:1.17 AS build-env
WORKDIR /data
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:3.13 AS final
WORKDIR /data
COPY --from=build-env /data/app .
ENV VERSION work02
EXPOSE 80
CMD ["./app"]

