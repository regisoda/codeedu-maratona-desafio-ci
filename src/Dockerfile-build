FROM golang:1.14.4-alpine as build
WORKDIR /app
COPY main.go app/
RUN go build -o ./app/main ./app/main.go

FROM scratch
COPY --from=build /app .
ENTRYPOINT [ "/app/main"]


