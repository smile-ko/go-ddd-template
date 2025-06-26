FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/

FROM scratch

COPY --from=builder /build/server .
COPY configs /configs

ENTRYPOINT ["/server"]
