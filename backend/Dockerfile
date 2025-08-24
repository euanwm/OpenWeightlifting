FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend

FROM alpine:3.16 AS final
WORKDIR /app

COPY --from=build /app/backend /app/backend

ENV HOSTNAME backend
EXPOSE 8080

ENTRYPOINT ["/app/backend"]
