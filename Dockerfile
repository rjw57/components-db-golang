########## Backend web server
FROM golang:1.22 AS backend-build
WORKDIR /usr/src/backend
COPY backend/go.mod backend/go.sum ./
RUN \
  --mount=type=cache,target=/root/.cache \
  go mod download
RUN \
  --mount=type=cache,target=/root/.cache \
  --mount=target=/usr/src \
  CGO_ENABLED=0 go build -o /go/bin/backend

FROM scratch AS backend
ENV GIN_MODE=release
COPY --from=backend-build /go/bin/backend /backend
ENTRYPOINT ["/backend"]
