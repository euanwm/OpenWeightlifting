FROM golang:1.20.2-alpine3.16 AS build
WORKDIR /work
COPY . .
RUN \
  --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg/mod \
  go build -o backend serverMain.go

FROM alpine:3.16 AS final
COPY --from=build /work/backend /bin/backend

WORKDIR /
ENTRYPOINT ["/bin/backend"]
