FROM golang:1.18.0-alpine3.15 AS build
WORKDIR /work
COPY . .
RUN \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  go build -o backend serverMain.go

FROM alpine:3.15 AS final
COPY --from=build /work/backend /bin/backend

WORKDIR /
ENTRYPOINT ["/bin/backend"]
