FROM golang:1.20 AS build-image

WORKDIR /app

COPY go.mod go.sum ./
RUN GOPROXY=direct go mod download

COPY /cmd/*.go ./
COPY /pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" main.go


FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build-image /app/main ./

ENTRYPOINT ["/app/main"]
