FROM golang:alpine as build

WORKDIR /app
COPY ./ /app

RUN CGO_ENABLED=0 go build -o check-spam cmd/check-spam/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=build /app/check-spam /app/
RUN chmod +x /app/check-spam
ENTRYPOINT /app/check-spam