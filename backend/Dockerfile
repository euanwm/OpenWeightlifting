FROM golang:1.21-alpine AS build
WORKDIR /work
COPY . .
RUN \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  go build -o backend

FROM alpine:3.16 AS final
COPY --from=build /work/backend /bin/backend

ENV HOSTNAME backend
EXPOSE 8080

WORKDIR /
ENTRYPOINT ["/bin/backend"]
